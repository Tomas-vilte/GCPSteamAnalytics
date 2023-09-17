package persistence

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
			db, _ = createClient()
		}
	})
	return db
}

func createClient() (*sqlx.DB, error) {
	dbUser := "root"
	dbPwd := "root"
	dbName := "steamAnalytics"
	instanceConnectionName := "gcpsteamanalytics:us-central1:my-db-instance"
	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sqlx.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}
