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
	data, err := dataFetcher.GetData()
	if err != nil {
		log.Println("Error al obtener datos:", err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	dba := &db.MySQLDatabase{}

	err = dba.Connect()
	if err != nil {
		log.Println("Error al conectar a Mysql:", err)
	}
	defer func() {
		if err := dba.Close(); err != nil {
			log.Println("Error al cerrar la conexión:", err)
		}
	}()

	for _, item := range data {
		newItem := handlers.Item{
			Appid: item.Appid,
			Name:  item.Name,
		}
		log.Println("Juego insertado en la tabla", newItem.Name)
		err := dba.Insert(newItem)
		if err != nil {
			log.Printf("Error al insertar el elemento: %v", err)
		}

	}
}
