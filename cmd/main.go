package main

import (
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	//storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(http.Client{})
	//gameProcessor := service.NewGameProcessor(storage, steamClient)
	data, err := steamClient.GetAppDetails(730)

	if err != nil {
		log.Printf("Hubo un error: %v", err)
		return
	}
	fmt.Println(data)
}
