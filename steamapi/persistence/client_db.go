package persistence

import (
	"database/sql"
	"sync"
)

var doOnce sync.Once
var db *sql.DB

func GetDB() *sql.DB {
	doOnce.Do(func() {
		if db == nil {
			db = createClient()
		}
	})
	return db
}

func createClient() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics")
	if err != nil {
		panic(err)
	}
	return db
}
