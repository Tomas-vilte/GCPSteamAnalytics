package persistence

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	"time"
)

type storage interface {
	GetAllFrom(limit int) ([]handlers.Item, error)
	Update(item handlers.Item) error
}

type Storage struct {
}

func (s Storage) GetAllFrom(limit int) ([]handlers.Item, error) {
	query := "SELECT * FROM game WHERE status = 'PENDING' AND valid = false ORDER BY id LIMIT ?"
	rows, err := GetDB().Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []handlers.Item
	for rows.Next() {
		var entity handlers.Item
		err := rows.Scan(&entities)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (s Storage) Update(item handlers.Item) error {
	query := "UPDATE game SET status = ?, valid = ?, updated_at = ? WHERE app_id = ?"
	_, err := GetDB().Exec(query, handlers.PROCESSED, item.IsValid, time.Now(), item.Appid)
	if err != nil {
		return err
	}
	return nil
}
