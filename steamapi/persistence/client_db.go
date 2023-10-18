package persistence

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"cloud.google.com/go/cloudsqlconn"
	config2 "github.com/Tomas-vilte/GCPSteamAnalytics/config"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db     *sqlx.DB
	doOnce sync.Once
)

func GetDB() *sqlx.DB {
	doOnce.Do(func() {
		if db == nil {
			var err error
			db = createClientLocal()
			if err != nil {
				log.Fatalf("Error al crear la conexión con la base de datos: %v", err)
			}
		}
		log.Println("Conexión creada con éxito")
	})
	return db
}

// Esta conexion sirve si no vas a usar servicios de gcp.
func createClientLocal() *sqlx.DB {
	db, err := sqlx.Open("mysql", "tomi:tomi@tcp(172.19.0.4:3307)/steamAnalytics?parseTime=true")
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
	return dbPool, nil
}
