package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

type UserConfig struct {
	Editors   []string
	Protected Protected
	Rewards   Rewards
}

type Rewards struct {
	BttvReward `json:"Bttv"`
}

type BttvReward struct {
	Title                             string `json:"title"`
	Prompt                            string `json:"prompt"`
	Cost                              int    `json:"cost"`
	Backgroundcolor                   string `json:"backgroundColor,omitempty"`
	IsMaxPerStreamEnabled             bool   `json:"isMaxPerStreamEnabled,omitempty"`
	MaxPerStream                      int    `json:"maxPerStream,omitempty"`
	IsUserInputRequired               bool   `json:"isUserInputRequired,omitempty"`
	IsMaxPerUserPerStreamEnabled      bool   `json:"isMaxPerUserPerStreamEnabled,omitempty"`
	MaxPerUserPerStream               int    `json:"maxPerUserPerStream,omitempty"`
	IsGlobalCooldownEnabled           bool   `json:"isGlobalCooldownEnabled,omitempty"`
	GlobalCooldownSeconds             int    `json:"globalCooldownSeconds,omitempty"`
	ShouldRedemptionsSkipRequestQueue bool   `json:"shouldRedemptionsSkipRequestQueue,omitempty"`
	Enabled                           bool
	ID                                string
}

type Protected struct {
	EditorFor []string
}

func createDefaultUserConfig() UserConfig {
	return UserConfig{
		Editors: []string{},
		Protected: Protected{
			EditorFor: []string{},
		},
		Rewards: Rewards{},
	}
}

func (s *Server) handleUserConfig(w http.ResponseWriter, r *http.Request) {
	ok, auth, _ := s.authenticate(r)
	if !ok {
		http.Error(w, "bad authentication", http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet {
		userConfig, err, _ := s.getUserConfig(auth.Data.UserID)
		if err != nil {
			http.Error(w, "can't recover config"+err.Error(), http.StatusBadRequest)
			return
		}

		managing := r.URL.Query().Get("managing")
		if managing != "" {
			ownerUserID, err := s.checkEditor(r, userConfig)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			managingUserConfig, err, _ := s.getUserConfig(ownerUserID)
			if err != nil {
				http.Error(w, "can't recover config"+err.Error(), http.StatusBadRequest)
				return
			}
			newEditorForNames := []string{}

			userData, err := s.helixClient.GetUsersByUserIds(userConfig.Protected.EditorFor)
			if err != nil {
				http.Error(w, "can't resolve editorFor in config "+err.Error(), http.StatusBadRequest)
			}
			for _, user := range userData {
				newEditorForNames = append(newEditorForNames, user.Login)
			}

			managingUserConfig.Editors = []string{}
			managingUserConfig.Protected.EditorFor = newEditorForNames
			writeJSON(w, managingUserConfig, http.StatusOK)
			return
		} else {
			newEditorNames := []string{}

			userData, err := s.helixClient.GetUsersByUserIds(userConfig.Editors)
			if err != nil {
				http.Error(w, "can't resolve editors in config "+err.Error(), http.StatusBadRequest)
			}
			for _, user := range userData {
				newEditorNames = append(newEditorNames, user.Login)
			}

			userConfig.Editors = newEditorNames

			newEditorForNames := []string{}

			userData, err = s.helixClient.GetUsersByUserIds(userConfig.Protected.EditorFor)
			if err != nil {
				http.Error(w, "can't resolve editorFor in config "+err.Error(), http.StatusBadRequest)
			}
			for _, user := range userData {
				newEditorForNames = append(newEditorForNames, user.Login)
			}

			userConfig.Protected.EditorFor = newEditorForNames
			writeJSON(w, userConfig, http.StatusOK)
			return
		}

	} else if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorf("Failed reading update body: %s", err)
			http.Error(w, "Failure saving body: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.processConfig(auth.Data.UserID, body, r)
		if err != nil {
			log.Errorf("failed processing config: %s", err)
			http.Error(w, "failed processing config: "+err.Error(), http.StatusBadRequest)
			return
		}

		writeJSON(w, "", http.StatusOK)
	} else if r.Method == http.MethodDelete {
		_, err := s.store.Client.HDel("userConfig", auth.Data.UserID).Result()
		if err != nil {
			log.Error(err)
			http.Error(w, "Failed deleting: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.unsubscribeChannelPoints(auth.Data.UserID)
		if err != nil {
			log.Error(err)
			http.Error(w, "Failed to unsubscribe"+err.Error(), http.StatusInternalServerError)
			return
		}

		writeJSON(w, "", http.StatusOK)
	}

}

func (s *Server) getUserConfig(userID string) (UserConfig, error, bool) {
	val, err := s.store.Client.HGet("userConfig", userID).Result()
	if err == redis.Nil || val == "" {
		return createDefaultUserConfig(), nil, true
	}
	if err != nil {
		return createDefaultUserConfig(), err, true
	}

	var userConfig UserConfig
	if err := json.Unmarshal([]byte(val), &userConfig); err != nil {
		return createDefaultUserConfig(), errors.New("can't find config"), true
	}

	return userConfig, nil, false
}

func (s *Server) checkEditor(r *http.Request, userConfig UserConfig) (string, error) {
	managing := r.URL.Query().Get("managing")

	if managing == "" {
		return "", nil
	}

	userData, err := s.helixClient.GetUsersByUsernames([]string{managing})
	if err != nil || len(userData) == 0 {
		return "", errors.New("can't find editor")
	}

	isEditor := false
	for _, editor := range userConfig.Protected.EditorFor {
		if editor == userData[managing].ID {
			isEditor = true
		}
	}

	if !isEditor {
		return "", errors.New("User is not editor")
	}

	return userData[managing].ID, nil
}

func (s *Server) processConfig(userID string, body []byte, r *http.Request) error {
	oldConfig, err, isNew := s.getUserConfig(userID)
	if err != nil {
		return err
	}

	ownerUserID, err := s.checkEditor(r, oldConfig)
	if err != nil {
		return err
	}

	if ownerUserID != "" {
		oldConfig, err, isNew = s.getUserConfig(ownerUserID)
		if err != nil {
			return err
		}
	}

	var newConfig UserConfig
	if err := json.Unmarshal(body, &newConfig); err != nil {
		return err
	}

	protected := oldConfig.Protected
	if protected.EditorFor == nil {
		protected.EditorFor = []string{}
	}

	configToSave := UserConfig{
		Editors:   oldConfig.Editors,
		Rewards:   newConfig.Rewards,
		Protected: protected,
	}

	if ownerUserID == "" {
		newEditorIds := []string{}

		userData, err := s.helixClient.GetUsersByUsernames(newConfig.Editors)
		if err != nil {
			return err
		}
		if len(newConfig.Editors) != len(userData) {
			return errors.New("Failed to find all editors")
		}

		for _, user := range userData {
			if user.ID == userID {
				return errors.New("You can't be your own editor")
			}

			newEditorIds = append(newEditorIds, user.ID)
		}

		configToSave.Editors = newEditorIds

		for _, editor := range oldConfig.Editors {
			err := s.removeEditorFor(editor, userID)
			if err != nil {
				return err
			}
		}

		for _, user := range userData {
			err := s.addEditorFor(user.ID, userID)
			if err != nil {
				return err
			}
		}
	}

	saveTarget := userID
	if ownerUserID != "" {
		saveTarget = ownerUserID
	}

	// the saving to twitch part
	if configToSave.Rewards.BttvReward.ID == "" {
		log.Info("creating reward")
		reward, err := s.createChannelPointReward(saveTarget, newConfig.Rewards.BttvReward)
		if err != nil {
			return err
		}

		newConfig.Rewards.BttvReward = reward
	} else {
		log.Info("updating reward")
		// update reward
	}

	js, err := json.Marshal(configToSave)
	if err != nil {
		return err
	}

	_, err = s.store.Client.HSet("userConfig", saveTarget, js).Result()
	if err != nil {
		return err
	}

	if isNew {
		log.Info("Created new config for: ", userID)
		s.subscribeChannelPoints(userID)
	}

	return nil
}

func writeJSON(w http.ResponseWriter, data interface{}, code int) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, err = w.Write(js)
	if err != nil {
		log.Errorf("Faile to writeJSON: %s", err)
	}
}

func (s *Server) addEditorFor(editorID, userID string) error {
	isNew := true

	val, err := s.store.Client.HGet("userConfig", editorID).Result()
	if err == redis.Nil {
		isNew = true
	} else if err != nil {
		return err
	}

	var userConfig UserConfig
	if !isNew {
		if err := json.Unmarshal([]byte(val), &userConfig); err != nil {
			return err
		}
	} else {
		userConfig = createDefaultUserConfig()
	}

	userConfig.Protected.EditorFor = append(userConfig.Protected.EditorFor, userID)

	js, err := json.Marshal(userConfig)
	if err != nil {
		return err
	}

	log.Infof("New Editor %s for user %s", editorID, userID)
	_, err = s.store.Client.HSet("userConfig", editorID, js).Result()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) removeEditorFor(editorID, userID string) error {
	isNew := true

	val, err := s.store.Client.HGet("userConfig", editorID).Result()
	if err == redis.Nil {
		isNew = true
	} else if err != nil {
		return err
	}

	var userConfig UserConfig
	if !isNew {
		if err := json.Unmarshal([]byte(val), &userConfig); err != nil {
			return err
		}
	} else {
		userConfig = createDefaultUserConfig()
	}

	newEditorFor := []string{}

	for _, editor := range userConfig.Protected.EditorFor {
		if userID != editor {
			newEditorFor = append(newEditorFor, editor)
		}
	}

	userConfig.Protected.EditorFor = newEditorFor

	js, err := json.Marshal(userConfig)
	if err != nil {
		return err
	}

	log.Infof("New Editor %s for user %s", editorID, userID)
	_, err = s.store.Client.HSet("userConfig", editorID, js).Result()
	if err != nil {
		return err
	}

	return nil
}
