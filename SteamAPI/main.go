package main

import (
	"fmt"
	"log"
	"steamAPI/api/db"
	"steamAPI/api/handlers"
)

func main() {
	//fmt.Println(config.GetCrendentials())
	//file := "/home/tomi/GCPSteamAnalytics/SteamAPI/api/data/hola2.txt"
	//fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))
	dataFetcher := &handlers.RealDataFetcher{}
	res, err := dataFetcher.GetData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}

	var dba db.MySQLDatabase

	err = dba.Connect()
	if err != nil {
		log.Fatal("Error al conectar a Mysql:", err)
	}
	defer func() {
		if err := dba.Close(); err != nil {
			log.Println("Error al cerrar la conexión:", err)
		}
	}()
}
