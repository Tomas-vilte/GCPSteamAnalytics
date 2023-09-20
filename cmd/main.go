package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Si vas a usarlo en local o en gcp acordate primero de ejecutar esto
	//data := &handlers.RealDataFetcher{}
	//db1 := &db.MySQLDatabase{}
	//fmt.Println(db.InsertData(data, db1))

	log.Printf("App started!")
	api.StartServer()
}
