package persistence

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"log"
	"time"
)

type StorageDB interface {
	GetAllFrom(limit int) ([]entity.Item, error)
	Update(item entity.Item) error
}

func NewStorage() StorageDB {
	return &storage{}
}

type storage struct{}

func (s storage) GetAllFrom(limit int) ([]entity.Item, error) {
	query := "SELECT app_id, name, status, valid, created_at, updated_at FROM game WHERE status = 'PENDING' AND valid = false ORDER BY id LIMIT ?"
	rows, err := GetDB().Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []entity.Item
	for rows.Next() {
		item := entity.Item{}
		err := rows.Scan(&item.Appid, &item.Name, &item.Status, &item.IsValid, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		entities = append(entities, item)
	}

	return entities, nil
}

func (s storage) Update(item entity.Item) error {
	query := "UPDATE game SET status = ?, valid = ?, updated_at = ? WHERE app_id = ?"
	_, err := GetDB().Exec(query, entity.PROCESSED, item.IsValid, time.Now(), item.Appid)
	if err != nil {
		log.Printf("Error al actualizar el estado del juego con appID %d: %v\n", item.Appid, err)
		return err
	}
	log.Printf("Estado actualizado del juego con appID %d\n", item.Appid)
	return nil
}
