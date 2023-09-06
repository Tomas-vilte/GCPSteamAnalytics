package persistence

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	db     *sqlx.DB
	doOnce sync.Once
)

func GetDB() *sqlx.DB {
	doOnce.Do(func() {
		if db == nil {
			db = createClient()
		}
	})
	return db
}

func createClient() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
