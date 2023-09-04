package main

import (
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/db"
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//log.Printf("App started!")
	//api.StartServer()
	fetcher := &handlers.RealDataFetcher{}
	db1 := &db.MySQLDatabase{}
	data := db.InsertData(fetcher, db1)
	fmt.Println(data)
}
