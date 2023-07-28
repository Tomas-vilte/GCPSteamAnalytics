package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"steamAPI/api/handlers"
)

const batchSize = 1000

type Database interface {
	Connect() error
	Close() error
	InsertBatch(items []handlers.Item) error
	InsertBatchData(items []handlers.Item) error
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

func (m *MySQLDatabase) Insert(item handlers.Item) error {
	_, err := m.db.Exec("INSERT INTO games (appid, name) VALUES (?, ?) ", item.Appid, item.Name)
	if err != nil {
		log.Printf("Error al insertar el elemento: %v", err)
		return err
	}
	return nil
}

func (m *MySQLDatabase) InsertBatch(items []handlers.Item) error {
	if len(items) == 0 {
		return nil
	}

	// Dividimos los datos en lotes
	numItems := len(items)
	numBatches := (numItems + batchSize - 1) / batchSize

	// Iteramos los lotes y realizamos la insercion por lotes
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

func (m *MySQLDatabase) InsertBatchData(items []handlers.Item) error {
	if len(items) == 0 {
		return nil
	}

	// Creamos la query para insertar los lotes
	query := "INSERT INTO games (appid, name) VALUES (?, ?)"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		log.Printf("Error al preparar la consulta: %v", err)
		return err
	}
	defer stmt.Close()

	// Ejecutamos la query en la base de datos utilizando transacciones
	tx, err := m.db.Begin()
	if err != nil {
		log.Printf("Error al iniciar la transaccion: %v", err)
		return err
	}
	tx.Rollback() // Si hay un error, se hace un rollback de la transaccion

	// Iteramos sobre los elementos y ejecutamos la query con los valores correspondientes
	for _, item := range items {
		_, err := tx.Stmt(stmt).Exec(item.Appid, item.Name)
		if err != nil {
			log.Printf("Error al insertar el lote de elementos: %v", err)
			return err
		}
	}
	// Hacemos commit de la transaccion una vez que se insertan todos los elementos
	if err = tx.Commit(); err != nil {
		log.Printf("Error al hacer commit de la transaccion: %v", err)
		return err
	}
	return nil
}
