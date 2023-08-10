package main

import (
	"database/sql"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics")
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Hubo un error al conectarse a la base de datos: %v", err)
		defer db.Close()
		return
	}

	steamAPI := &steamapi.SteamAPI{DB: db}

	gameDetails := steamAPI.ExtractAndSaveLimitedGameDetails(20)
	fmt.Println(gameDetails)
}
