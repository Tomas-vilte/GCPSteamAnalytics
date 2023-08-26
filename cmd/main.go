package main

import (
	"context"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	gameProcessor := service.NewGameProcessor(nil, nil)
	err := gameProcessor.RunProcessData(context.Background(), 20)

	if err != nil {
		log.Printf("Hubo un error: %v", err)
		return
	}
}
