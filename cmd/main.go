package main

import (
	"database/sql"
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

	lastProcessedAppID, err := steamAPI.LoadLastProcessedAppid()
	if err != nil {
		log.Printf("Error al cargar el último appID procesado: %v", err)
		return
	}

	// Cargar SteamAppIDs previamente procesados
	appIDs, err := steamAPI.GetAllAppIDs(lastProcessedAppID)
	if err != nil {
		log.Printf("Error al obtener los appIDs: %v", err)
		return
	}
	data, err := steamAPI.GetSteamData(appIDs, 10)
	if err != nil {
		log.Printf("Error al obtener los datos de Steam: %v", err)
		return
	}

	err = steamapi.SaveToCSV(data, "../data/dataDetails.csv")
	if err != nil {
		log.Printf("Error al guardar el CSV: %v", err)
		return
	}

}
