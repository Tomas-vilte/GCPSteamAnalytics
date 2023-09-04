package entity

import (
	"database/sql"
	"time"
)

type status string

const (
	PENDING   status = "PENDING"
	PROCESSED status = "PROCESSED"
)

type Item struct {
	Appid     int64         `json:"appid"`
	Name      string        `json:"name"`
	Status    status        `json:"status"`
	IsValid   bool          `json:"valid"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *sql.NullTime `json:"updated_at,omitempty"`
}
