package db

import (
	"database/sql"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const batchSize = 1000

type Database interface {
	Connect() error
	Close() error
	InsertBatch(items []handlers.Item) error
	InsertBatchData(items []handlers.Item) error
	GetAppIDs() ([]int, error)
}

type MySQLDatabase struct {
	db *sql.DB
}

func (m *MySQLDatabase) Connect() error {
	var err error
	m.db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics")
	if err != nil {
		return err
	}

	err = m.db.Ping()
	if err != nil {
		log.Printf("Hubo un error al conectarse a la base de datos: %v", err)
		m.db.Close()
		return err
	}

	fmt.Println("Conexión exitosa a MySQL")
	return nil
}

func (m *MySQLDatabase) Close() error {
	return m.db.Close()
}

// InsertBatch realiza la inserción de datos en la base de datos por lotes.
// Divide los datos ingresados en lotes más pequeñas y llama a la función InsertBatchData para cada lote.
// Si ocurre algún problema durante la inserción en lotes, la función devuelve un error.
func (m *MySQLDatabase) InsertBatch(items []handlers.Item) error {
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
func (m *MySQLDatabase) InsertBatchData(items []handlers.Item) error {
	if len(items) == 0 {
		return nil
	}

	// Creamos la consulta para la insercion en lotes
	query := "INSERT INTO gamesdetails (appid, name) VALUES "
	var vals []interface{}
	for i, item := range items {
		query += "(?, ?)"
		vals = append(vals, item.Appid, item.Name)
		if i < len(items)-1 {
			query += ", "
		}
	}

	// Ejecutamos la consulta en la base de datos
	_, err := m.db.Exec(query, vals...)
	if err != nil {
		log.Printf("Error al insertar el lote de elementos: %v", err)
		return err
	}

	return nil
}

// GetAppIDs obtiene todos los appid almacenados en la base de datos MySQL.
func (m *MySQLDatabase) GetAppIDs() ([]int, error) {
	rows, err := m.db.Query("SELECT appid FROM games")
	if err != nil {
		log.Printf("Error al obtener los appid desde la base de datos: %v", err)
		return nil, err
	}
	defer rows.Close()

	var appids []int
	for rows.Next() {
		var appid int
		if err := rows.Scan(&appid); err != nil {
			log.Printf("Error al escanear el appid desde la base de datos: %v", err)
			return nil, err
		}
		appids = append(appids, appid)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al obtener los appid desde la base de datos: %v", err)
		return nil, err
	}

	return appids, nil
}
