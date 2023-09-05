package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	log.Printf("App started!")
	api.StartServer()

	//test := persistence.NewStorage()
	//fmt.Println(test.GetGameDetails(10))
}
