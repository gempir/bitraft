package main

import (
	"flag"

	"github.com/gempir/bitraft/pkg/config"
	"github.com/gempir/bitraft/pkg/helix"
	"github.com/gempir/bitraft/pkg/store"
	"github.com/gempir/bitraft/services/api/emotechief"
	"github.com/gempir/bitraft/services/api/server"
)

func main() {
	configFile := flag.String("config", "../../config.json", "json config file")
	flag.Parse()

	cfg := config.NewConfig(*configFile)
	rStore := store.NewRedis()
	db := store.NewDatabase(cfg.SqliteDatabase)

	helixClient := helix.NewClient(cfg.ClientID, cfg.ClientSecret, "")
	helixUserClient := helix.NewClient(cfg.ClientID, cfg.ClientSecret, cfg.ApiBaseUrl+"/api/callback")
	go helixUserClient.StartRefreshTokenRoutine()
	go helixClient.StartRefreshTokenRoutine()

	emoteChief := emotechief.NewEmoteChief(rStore, cfg, db)
	server := server.NewServer(cfg, helixClient, helixUserClient, rStore, db, emoteChief)

	server.Start()
}
