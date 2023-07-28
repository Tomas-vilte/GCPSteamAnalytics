package db

import (
	"log"
	"steamAPI/api/handlers"
)

// InsertData realiza la inserción de datos en la base de datos utilizando la carga por lotes.
func InsertData(fetcher handlers.DataFetcher, database Database) error {
	// Obtener los datos desde la API
	items, err := fetcher.GetData()
	if err != nil {
		log.Printf("Error al obtener los datos desde la API: %v", err)
		return err
	}

	// Crear la instancia de la base de datos y conectar
	err = database.Connect()
	if err != nil {
		log.Printf("Error al conectar a la base de datos: %v", err)
		return err
	}
	defer database.Close()

	// Insertar los datos en la base de datos en lotes
	err = database.InsertBatch(items)
	if err != nil {
		log.Printf("Error al insertar los datos en lotes: %v", err)
		return err
	}

	log.Println("Los datos se han cargado correctamente en la base de datos.")
	return nil
}
