package main

import (
	"fmt"
	"steamAPI/api/funcgcp"

	"github.com/gin-gonic/gin"
)

// const endpointURL = "https://us-central1-gcpsteamanalytics.cloudfunctions.net"

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
	r := gin.Default()
	r.POST("/dbgames", funcgcp.ProcessSteamDataAndSaveToStorage)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("error:", err)
	}

}
