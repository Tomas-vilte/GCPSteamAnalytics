package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	//log.Printf("App started!")
	//api.StartServer()
	db, err := sql.Open("mysql", "my-db-user:root@tcp(34.42.224.196:3306)/steamAnalytics")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Lee el contenido del archivo schema.sql.
	sqlFile, err := os.ReadFile("../db/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Ejecuta el contenido del archivo en la base de datos.
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tablas creadas exitosamente")
}
