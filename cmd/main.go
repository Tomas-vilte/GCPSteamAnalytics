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
	//fetcher := &handlers.RealDataFetcher{}
	//dba := &db2.MySQLDatabase{}
	//
	//err = db2.InsertData(fetcher, dba)
	//if err != nil {
	//	log.Printf("Error al insertar los datos: %v", err)
	//}

	steamAPI := &steamapi.SteamAPI{DB: db, Client: &http.Client{}}
	err = steamapi.RunProcessData(steamAPI, 20)
	if err != nil {
		log.Printf("Hubo un error: %v", err)
		return
	}

	//reviewAPI := &steamapi.SteamReviewAPI{Client: &http.Client{}}
	//appids, err := utils.ReadAppIDsFromCSV("../data/gamesDetails1.csv")
	//if err != nil {
	//	log.Printf("Error al leer los appids: %v", err)
	//	return
	//}
	//
	//for _, appid := range appids {
	//	// Procesar reseñas para appid
	//	reviews, err := reviewAPI.GetReviews(appid)
	//	if err != nil {
	//		log.Printf("Error al obtener la reseña: %v", err)
	//		continue
	//	}
	//	err = reviewAPI.SaveReviewsToCSV(appid, reviews, "../data/reviewsNegative1.csv")
	//	if err != nil {
	//		log.Printf("Error al guardar las reseñas en CSV: %v", err)
	//	}
	//
	//}

}
