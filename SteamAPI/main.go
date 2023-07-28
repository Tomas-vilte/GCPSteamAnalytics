package main

import (
	"log"
	"steamAPI/api/db"
	"steamAPI/api/handlers"
)

func main() {
	//fmt.Println(config.GetCrendentials())
	//file := "/home/tomi/GCPSteamAnalytics/SteamAPI/api/data/hola2.txt"
	//fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))
	fetcher := &handlers.RealDataFetcher{}
	database := &db.MySQLDatabase{}

	err := db.InsertData(fetcher, database)
	if err != nil {
		log.Fatalf("Error al cargar los datos en la base de datos: %v", err)
	}
}
