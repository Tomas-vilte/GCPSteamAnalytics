package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Database interface {
	Connect() error
	Close() error
}

type MySQLDatabase struct {
	db *sql.DB
}

func (m *MySQLDatabase) Connect() error {
	var err error
	m.db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/goweb")
	if err != nil {
		log.Fatalf("Hubo un error al conectarse al Mysql: %v", err)
		return err
	}

	err = m.db.Ping()
	if err != nil {
		log.Fatalf("Error al cerrar la conexion: %v", err)
		m.db.Close()
		return err
	}

	fmt.Println("Conexión exitosa a MySQL")
	return nil
}

func (m *MySQLDatabase) Close() error {
	return m.db.Close()
}
