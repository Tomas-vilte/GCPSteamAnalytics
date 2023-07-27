package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"steamAPI/api/handlers"
)

type Database interface {
	Connect() error
	Close() error
	Insert(item handlers.Item) error
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
