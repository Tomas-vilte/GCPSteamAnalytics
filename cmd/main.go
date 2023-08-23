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
	err = steamapi.RunProcessData(steamAPI, 100)
	if err != nil {
		return
	}

	//reviewAPI := &steamapi.SteamReviewAPI{Client: &http.Client{}}
	//appids, err := utils.ReadAppIDsFromCSV("/home/tomi/GCPSteamAnalytics/data/gamesDetails.csv")
	//if err != nil {
	//	log.Printf("Error al leer los appids: %v", err)
	//}
	//
	//for _, appid := range appids {
	//	reviews, err := reviewAPI.GetReviews(appid)
	//	if err != nil {
	//		log.Printf("Error al obtener la reseña: %v", err)
	//		continue
	//	}
	//	err = reviewAPI.SaveReviewsToCSV(appid, reviews, "../data/reviewsPositive.csv")
	//}
}
