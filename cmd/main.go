package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Si vas a usarlo en local o en gcp acordate primero de ejecutar esto
	//data := &handlers.RealDataFetcher{}
	//db := &db2.MySQLDatabase{}
	//fmt.Println(db2.InsertData(data, db))

	log.Printf("App started!")
	api.StartServer()
}
