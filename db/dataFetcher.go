package db

import (
	"log"

	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
)

// InsertData realiza la inserción de datos en la base de datos utilizando la carga por lotes.
func InsertData(fetcher handlers.DataFetcher, database Database) error {
	// Obtenemos los datos de la API
	items, err := fetcher.GetData()
	if err != nil {
		log.Printf("Error al obtener los datos desde la API: %v", err)
		return err
	}

	// Insertamos los datos a la base de datos
	err = database.InsertBatch(items)
	if err != nil {
		log.Printf("Error al insertar los datos en lotes: %v", err)
		return err
	}

	log.Println("Los datos se han cargado correctamente en la base de datos.")
	return nil
}
