package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	log.Printf("App started!")
	api.StartServer()
}
