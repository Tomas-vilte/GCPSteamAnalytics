package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	db2 "github.com/Tomas-vilte/GCPSteamAnalytics/db"
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	log.Printf("App started!")
	api.StartServer()
	data := handlers.RealDataFetcher{}
	db := db2.MySQLDatabase{}
	db2.InsertData(&data, &db)

}
