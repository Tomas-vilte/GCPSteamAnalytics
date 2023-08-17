package main

import (
	"database/sql"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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
	steamAPI := &steamapi.SteamAPI{DB: db, Client: &http.Client{}}

	err = steamapi.RunProcessData(steamAPI, 5)
	if err != nil {

	}
}
