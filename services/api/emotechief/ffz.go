package emotechief

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type ffzEmoteResponse struct {
	Emote struct {
		ID       int         `json:"id"`
		Name     string      `json:"name"`
		Height   int         `json:"height"`
		Width    int         `json:"width"`
		Public   bool        `json:"public"`
		Hidden   bool        `json:"hidden"`
		Modifier bool        `json:"modifier"`
		Offset   interface{} `json:"offset"`
		Margins  interface{} `json:"margins"`
		CSS      interface{} `json:"css"`
		Owner    struct {
			ID          int    `json:"_id"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
		} `json:"owner"`
		Urls struct {
			Num1 string `json:"1"`
		} `json:"urls"`
		Status      int       `json:"status"`
		UsageCount  int       `json:"usage_count"`
		CreatedAt   time.Time `json:"created_at"`
		LastUpdated time.Time `json:"last_updated"`
	} `json:"emote"`
}

func (e *EmoteChief) SetFfzEmote(channelUserID, emoteId, channel string) (addedEmote *ffzEmoteResponse, removedEmote *ffzEmoteResponse, err error) {
	addedEmote, err = getFfzEmote(emoteId)
	if err != nil {
		return nil, nil, err
	}

	if !addedEmote.Emote.Public {
		return nil, nil, errors.New("emote is not public")
	}

	return nil, nil, nil
}

func getFfzEmote(emoteID string) (*ffzEmoteResponse, error) {
	if emoteID == "" {
		return nil, errors.New("emoteId invalid")
	}

	response, err := http.Get("https://api.frankerfacez.com/v1/emote/" + emoteID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if response.StatusCode <= 100 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("Bad bttv response: %d", response.StatusCode)
	}

	var emoteResponse ffzEmoteResponse
	err = json.NewDecoder(response.Body).Decode(&emoteResponse)
	if err != nil {
		return nil, err
	}

	return &emoteResponse, nil
}
