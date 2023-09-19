package db

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"log"

	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	_ "github.com/go-sql-driver/mysql"
)

const batchSize = 1000

type Database interface {
	InsertBatch(items []entity.Item) error
	InsertBatchData(items []entity.Item) error
}

type MySQLDatabase struct{}

// InsertBatch realiza la inserción de datos en la base de datos por lotes.
// Divide los datos ingresados en lotes más pequeñas y llama a la función InsertBatchData para cada lote.
// Si ocurre algún problema durante la inserción en lotes, la función devuelve un error.
func (m *MySQLDatabase) InsertBatch(items []entity.Item) error {
	// Dividimos los datos en lotes
	numItems := len(items)
	numBatches := (numItems + batchSize - 1) / batchSize

	// Iteramos los lotes y realizamos la inserción por lotes
	for i := 0; i < numBatches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize
		if end > numItems {
			end = numItems
		}
		batchData := items[start:end]

		err := m.InsertBatchData(batchData)
		if err != nil {
			log.Printf("Error al insertar el lote de elementos: %v", err)
			return err
		}
	}

	return nil
}

// InsertBatchData realiza la inserción en la base de datos de una tanda de elementos.
// La función recibe una lista de elementos (items) y construye una consulta SQL para insertarlos en la tabla "games".
// Utiliza marcadores de posición (?) para evitar inyecciones de SQL y luego ejecuta la consulta en la base de datos.
// Los valores de los elementos se proporcionan como argumentos para la consulta utilizando el operador "..." para desempaquetar el slice de valores.
func (m *MySQLDatabase) InsertBatchData(items []entity.Item) error {
	if len(items) == 0 {
		return nil
	}

	// Creamos la consulta para la inserción en lotes
	query := "INSERT INTO game (app_id, name, status, valid, created_at) VALUES "
	var vals []interface{}
	for i, item := range items {
		query += "(?, ?, ?, ?, ?)"
		vals = append(vals, item.Appid, item.Name, item.Status, item.IsValid, item.CreatedAt)
		if i < len(items)-1 {
			query += ", "
		}
	}

	// Ejecutamos la consulta en la base de datos
	_, err := persistence.GetDB().Exec(query, vals...)
	if err != nil {
		log.Printf("Error al insertar el lote de elementos: %v", err)
		return err
	}

	return nil
}
