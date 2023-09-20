package persistence

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"fmt"
	config2 "github.com/Tomas-vilte/GCPSteamAnalytics/config"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net"
	"sync"
)

var (
	db     *sqlx.DB
	doOnce sync.Once
)

func GetDB() *sqlx.DB {
	doOnce.Do(func() {
		if db == nil {
			db, _ = createClientInGCP() // Aca podes cambiarlo a createClientLocal() si no pensas usarlo en gcp
		}
	})
	return db
}

// Esta conexion sirve si no vas a usar servicios de gcp.
func createClientLocal() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

// Conexion con Google Cloud SQL
func createClientInGCP() (*sqlx.DB, error) {
	// !ACORDATE DE CONFIGURAR LAS VARIABLES DE ENTORNO EN GCP!
	config := config2.LoadEnvVariables()
	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, config.InstanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		config.DBUser, config.DBPass, config.DBName)

	dbPool, err := sqlx.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	log.Println("Conexion creada con exito")
	return dbPool, nil
}
