package main

import (
	"net/http"
	"steamAPI/api/functions"
)

const endpointURL = "https://us-central1-gcpsteamanalytics.cloudfunctions.net"

func main() {
	//fmt.Println(config.GetCrendentials())
	//file := "/home/tomi/GCPSteamAnalytics/steamAPI/api/data/hola2.txt"
	//fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))
	//fetcher := &handlers.RealDataFetcher{}
	//database := &db.MySQLDatabase{}
	//
	//err := db.InsertData(fetcher, database)
	//if err != nil {
	//	log.Fatalf("Error al cargar los datos en la base de datos: %v", err)
	//}
	http.HandleFunc("/dbgames", functions.ProcessSteamDataAndSaveToStorage)

}
