package entity

import "database/sql"

type status string

const (
	PENDING   status = "PENDING"
	PROCESSED status = "PROCESSED"
)

type Item struct {
	Appid     int64         `json:"app_id"`
	Name      string        `json:"name"`
	Status    status        `json:"status"`
	IsValid   bool          `json:"valid"`
	CreatedAt *sql.NullTime `json:"created_at"`
	UpdatedAt *sql.NullTime `json:"updated_at,omitempty"`
}
