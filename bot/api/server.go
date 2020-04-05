package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gempir/spamchamp/bot/config"
	"github.com/gempir/spamchamp/bot/helix"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

// Server api server
type Server struct {
	cfg            *config.Config
	broadcastQueue chan BroadcastMessage
	helixClient    *helix.Client
}

type BroadcastMessage struct {
	ChannelStats   []ChannelStat `json:"channelStats"`
	ActiveChannels int           `json:"activeChannels"`
}

type ChannelStat struct {
	ID    string `json:"id"`
	Msgps int64  `json:"msgps"`
}

// NewServer create api Server
func NewServer(cfg *config.Config, helixClient *helix.Client, broadcastQueue chan BroadcastMessage) Server {
	return Server{
		cfg:            cfg,
		broadcastQueue: broadcastQueue,
		helixClient:    helixClient,
	}
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var upgrader = websocket.Upgrader{}

func (s *Server) Start() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	go s.handleMessages()
	http.HandleFunc("/api/ws", s.handleConnections)
	http.Handle("/api/channel", corsHandler(http.HandlerFunc(s.handleChannel)))

	err := http.ListenAndServe(":8035", nil)
	log.Info("[api] listening on port :8035")
	if err != nil {
		log.Fatal("[api] listenAndServe: ", err)
	}
}

func (s *Server) handleChannel(w http.ResponseWriter, r *http.Request) {
	channelIDs := strings.Split(r.URL.Query().Get("channelids"), ",")

	users, err := s.helixClient.GetUsersByUserIds(channelIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(users, http.StatusOK, w, r)
}

func (s *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true
	// Make sure we close the connection when the function returns
	defer ws.Close()

	for {
		var msg BroadcastMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Infof("[api] error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		s.broadcastQueue <- msg
	}
}

func (s *Server) handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-s.broadcastQueue
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Errorf("[api] error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func writeJSON(data interface{}, code int, w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}

func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			h.ServeHTTP(w, r)
		}
	})
}
