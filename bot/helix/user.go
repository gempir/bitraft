package helix

import (
	"fmt"
	"net/http"
	"strings"

	helixClient "github.com/nicklaw5/helix"
	log "github.com/sirupsen/logrus"
)

// Client wrapper for helix
type Client struct {
	clientID   string
	client     *helixClient.Client
	httpClient *http.Client
}

var (
	userCacheByID       map[string]*UserData
	userCacheByUsername map[string]*UserData
)

func init() {
	userCacheByID = map[string]*UserData{}
	userCacheByUsername = map[string]*UserData{}
}

// NewClient Create helix client
func NewClient(clientID string) Client {
	client, err := helixClient.NewClient(&helixClient.Options{
		ClientID: clientID,
	})
	if err != nil {
		panic(err)
	}

	return Client{
		clientID:   clientID,
		client:     client,
		httpClient: &http.Client{},
	}
}

type userResponse struct {
	Data []UserData `json:"data"`
}

// UserData exported data from twitch
type UserData struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

// GetUsersByUserIds receive userData for given ids
func (c *Client) GetUsersByUserIds(userIDs []string) (map[string]UserData, error) {
	var filteredUserIDs []string

	for _, id := range userIDs {
		if _, ok := userCacheByID[id]; !ok {
			filteredUserIDs = append(filteredUserIDs, id)
		}
	}

	if len(filteredUserIDs) > 0 {
		var chunks [][]string
		for i := 0; i < len(filteredUserIDs); i += 99 {
			end := i + 99

			if end > len(filteredUserIDs) {
				end = len(filteredUserIDs)
			}

			chunks = append(chunks, filteredUserIDs[i:end])
		}

		for _, chunk := range chunks {
			resp, err := c.client.GetUsers(&helixClient.UsersParams{
				IDs: chunk,
			})
			if err != nil {
				return make(map[string]UserData), err
			}
			log.Infof("[helix] %d GetUsersByUserIds %v", resp.StatusCode, chunk)
			if resp.StatusCode > http.StatusMultipleChoices {
				return make(map[string]UserData), fmt.Errorf("bad helix response: %v", resp.ErrorMessage)
			}

			for _, user := range resp.Data.Users {
				data := &UserData{
					ID:              user.ID,
					Login:           user.Login,
					DisplayName:     user.Login,
					Type:            user.Type,
					BroadcasterType: user.BroadcasterType,
					Description:     user.Description,
					ProfileImageURL: user.ProfileImageURL,
					OfflineImageURL: user.OfflineImageURL,
					ViewCount:       user.ViewCount,
					Email:           user.Email,
				}
				userCacheByID[user.ID] = data
				userCacheByUsername[user.Login] = data
			}
		}
	}

	result := make(map[string]UserData)

	for _, id := range userIDs {
		result[id] = *userCacheByID[id]
	}

	return result, nil
}

// GetUsersByUsernames fetches userdata from helix
func (c *Client) GetUsersByUsernames(usernames []string) (map[string]UserData, error) {
	var filteredUsernames []string

	for _, username := range usernames {
		if _, ok := userCacheByUsername[strings.ToLower(username)]; !ok {
			filteredUsernames = append(filteredUsernames, strings.ToLower(username))
		}
	}

	if len(filteredUsernames) > 0 {
		resp, err := c.client.GetUsers(&helixClient.UsersParams{
			Logins: filteredUsernames,
		})
		if err != nil {
			return map[string]UserData{}, err
		}

		log.Infof("[helix] %d GetUsersByUsernames %v", resp.StatusCode, filteredUsernames)

		for _, user := range resp.Data.Users {
			data := &UserData{
				ID:              user.ID,
				Login:           user.Login,
				DisplayName:     user.Login,
				Type:            user.Type,
				BroadcasterType: user.BroadcasterType,
				Description:     user.Description,
				ProfileImageURL: user.ProfileImageURL,
				OfflineImageURL: user.OfflineImageURL,
				ViewCount:       user.ViewCount,
				Email:           user.Email,
			}
			userCacheByID[user.ID] = data
			userCacheByUsername[user.Login] = data
		}
	}

	result := make(map[string]UserData)

	for _, username := range usernames {
		result[strings.ToLower(username)] = *userCacheByUsername[strings.ToLower(username)]
	}

	return result, nil
}
