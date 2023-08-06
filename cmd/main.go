package main

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/db"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"log"
)

func main() {
	dba := &db.MySQLDatabase{}
	err := dba.Connect()
	if err != nil {
		log.Printf("Error al conectar a la bd: %v", err)
	}
	defer dba.Close()

	steamAPI := &steamapi.SteamAPI{
		Dba: dba,
	}

	gameDetails, err := steamAPI.ExtractGameDetails()
	if err != nil {
		log.Printf("Error al obtener los detalles de los juegos: %v", err)
	}
	for _, game := range gameDetails {
		log.Printf("ID de cada juego %d:\n\n\n", game.SteamAppid)
		log.Printf("Nombre de cada juego: %s\n\n", game.NameGame)
		log.Printf("Descripcion de cada juego: %s\n\n", game.ShortDescription)
	}
}
