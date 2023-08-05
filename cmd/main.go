package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/db"
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	"log"
)

func main() {

	fetcher := &handlers.RealDataFetcher{}
	database := &db.MySQLDatabase{}

	err := db.InsertData(fetcher, database)
	if err != nil {
		log.Fatalf("Error al cargar los datos en la base de datos: %v", err)
	}
}
