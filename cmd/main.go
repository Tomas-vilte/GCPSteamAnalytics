package main

import (
	"database/sql"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
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
	err = steamapi.RunProcessData(steamAPI, 2000)
	if err != nil {
		log.Printf("Hubo un error: %v", err)
		return
	}

	reviewAPI := &steamapi.SteamReviewAPI{Client: &http.Client{}}
	appids, err := utils.ReadAppIDsFromCSV("/home/tomi/GCPSteamAnalytics/data/gamesDetails.csv")
	if err != nil {
		log.Printf("Error al leer los appids: %v", err)
		return
	}

	// Procesar reseñas desde el índice startIndex
	for _, appid := range appids {
		// Procesar reseñas para appid
		reviews, err := reviewAPI.GetReviews(appid)
		if err != nil {
			log.Printf("Error al obtener la reseña: %v", err)
			continue
		}
		err = reviewAPI.SaveReviewsToCSV(appid, reviews, "../data/reviewsNegative.csv")
		if err != nil {
			log.Printf("Error al guardar las reseñas en CSV: %v", err)
		}

	}

}
