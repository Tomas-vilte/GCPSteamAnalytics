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
	// Cargar SteamAppIDs previamente procesados
	appIDs, err := steamAPI.GetAllAppIDs()
	if err != nil {
		log.Printf("Error al obtener los appIDs: %v", err)
		return
	}
	data, err := steamapi.GetSteamData(appIDs, 50)
	if err != nil {
		log.Printf("Error al obtener los datos de Steam: %v", err)
		return
	}

	// Imprimir información de los juegos obtenidos
	for _, game := range data {
		log.Printf("Juego: %s, AppID: %d\n", game.Name, game.SteamAppid)
	}

	err = steamapi.SaveToCSV(data, "Output.csv")
	if err != nil {
		log.Printf("error saving data to CSV: %v", err)
		return
	}

}
