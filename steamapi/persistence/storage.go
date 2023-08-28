package persistence

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"time"
)

type Storage interface {
	GetAllFrom(limit int) ([]entity.Item, error)
	Update(item entity.Item) error
}

func NewStorage() Storage {
	return &storage{}
}

type storage struct {
}

func (s storage) GetAllFrom(limit int) ([]entity.Item, error) {
	query := "SELECT app_id, name, status, valid, created_at, updated_at FROM game WHERE status = 'PENDING' AND valid = false ORDER BY id LIMIT ?"
	rows, err := GetDB().Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []entity.Item
	for rows.Next() {
		entity := entity.Item{}
		err := rows.Scan(&entity.Appid, &entity.Name, &entity.Status, &entity.IsValid, &entity.CreatedAt, &entity.UpdatedAt)
		if err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (s storage) Update(item entity.Item) error {
	query := "UPDATE game SET status = ?, valid = ?, updated_at = ? WHERE app_id = ?"
	_, err := GetDB().Exec(query, entity.PROCESSED, item.IsValid, time.Now(), item.Appid)
	if err != nil {
		return err
	}
	return nil
}
