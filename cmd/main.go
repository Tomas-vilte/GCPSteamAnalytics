package main

import (
	"context"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(http.Client{})
	gameProcessor := service.NewGameProcessor(storage, steamClient)

	err := gameProcessor.RunProcessData(context.Background(), 30)

	if err != nil {
		log.Printf("Hubo un error: %v", err)
		return
	}
}
