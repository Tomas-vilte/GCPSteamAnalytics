package main

import (
	"database/sql"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//dba := &db.MySQLDatabase{}
	//err := dba.Connect()
	//if err != nil {
	//	log.Printf("Error al conectar a la bd: %v", err)
	//}
	//defer dba.Close()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics")
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Hubo un error al conectarse a la base de datos: %v", err)
		db.Close()
		return
	}

	steamAPI := &steamapi.SteamAPI{DB: db}

	gameDetails := steamAPI.ExtractAndSaveLimitedGameDetails(1000)
	fmt.Println(gameDetails)
}
