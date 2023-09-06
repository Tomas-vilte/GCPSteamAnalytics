package persistence

import (
	"database/sql"
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
	"log"
	"strings"
	"time"
)

type StorageDB interface {
	GetAllFrom(limit int) ([]entity.Item, error)
	Update(item entity.Item) error
	SaveGameDetails(dataProcessed []model.AppDetails) error
	GetGameDetails(id int) (*entity.GameDetails, error)
	GetAllByAppID(appID int) ([]entity.Item, error)
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

func (s storage) GetAllByAppID(appID int) ([]entity.Item, error) {
	query := "SELECT app_id, name, status, valid, created_at, updated_at FROM game WHERE app_id = ?"
	rows, err := GetDB().Query(query, appID)
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
	// Verificar si el juego existe en la tabla game.
	var count int
	err := GetDB().QueryRow("SELECT COUNT(*) FROM game WHERE app_id = ?", item.Appid).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar si el juego existe: %v\n", err)
		return err
	}

	if count == 0 {
		// El juego no existe en la tabla game, así que agrégalo primero.
		_, err := GetDB().Exec("INSERT INTO game (app_id, name, status, valid, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
			item.Appid, item.Name, entity.PROCESSED, item.IsValid, time.Now(), time.Now())
		if err != nil {
			log.Printf("Error al insertar el juego en la tabla game: %v\n", err)
			return err
		}
	}

	// Luego, realizar la actualización.
	query := "UPDATE game SET status = ?, valid = ?, updated_at = ? WHERE app_id = ?"
	_, err = GetDB().Exec(query, entity.PROCESSED, item.IsValid, time.Now(), item.Appid)
	if err != nil {
		log.Printf("Error al actualizar el estado del juego con appID %d: %v\n", item.Appid, err)
		return err
	}

	log.Printf("Estado actualizado del juego con appID %d\n", item.Appid)
	return nil
}

func (s storage) SaveGameDetails(dataProcessed []model.AppDetails) error {
	for _, appDetail := range dataProcessed {
		query := `
            INSERT INTO games_details (
                app_id, 
                name, 
                description, 
                type, 
                publishers, 
                developers, 
                is_free, 
                interface_languages, 
                fullAudio_languages, 
                subtitles_languages, 
                windows, 
                mac, 
                linux, 
                release_date, 
                coming_soon, 
                currency, 
                discount_percent, 
                initial_formatted, 
                final_formatted
            )
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE
                name = VALUES(name),
                description = VALUES(description),
                type = VALUES(type),
                publishers = VALUES(publishers),
                developers = VALUES(developers),
                is_free = VALUES(is_free),
                interface_languages = VALUES(interface_languages),
                fullAudio_languages = VALUES(fullAudio_languages),
                subtitles_languages = VALUES(subtitles_languages),
                windows = VALUES(windows),
                mac = VALUES(mac),
                linux = VALUES(linux),
                release_date = VALUES(release_date),
                coming_soon = VALUES(coming_soon),
                currency = VALUES(currency),
                discount_percent = VALUES(discount_percent),
                initial_formatted = VALUES(initial_formatted),
                final_formatted = VALUES(final_formatted)
        `
		_, err := GetDB().Exec(
			query,
			appDetail.SteamAppid,
			appDetail.Name,
			appDetail.Description,
			appDetail.Type,
			strings.Join(appDetail.Publishers, ", "),
			strings.Join(appDetail.Developers, ", "),
			appDetail.IsFree,
			utils.GetSupportedLanguagesString(appDetail.SupportedLanguages["interface"]),
			utils.GetSupportedLanguagesString(appDetail.SupportedLanguages["full_audio"]),
			utils.GetSupportedLanguagesString(appDetail.SupportedLanguages["subtitles"]),
			appDetail.Platforms.Windows,
			appDetail.Platforms.Mac,
			appDetail.Platforms.Linux,
			appDetail.ReleaseDate.Date,
			appDetail.ReleaseDate.ComingSoon,
			appDetail.PriceOverview.Currency,
			appDetail.PriceOverview.DiscountPercent,
			appDetail.PriceOverview.Initial,
			appDetail.PriceOverview.FinalFormatted,
		)
		if err != nil {
			log.Printf("Hubo un error al guardar los juegos: %v\n", err)
			return err
		}
	}

	return nil
}

func (s storage) GetGameDetails(gameID int) (*entity.GameDetails, error) {
	query := `SELECT * FROM games_details WHERE app_id = ?`

	var gameDetails entity.GameDetails
	err := GetDB().QueryRowx(query, gameID).StructScan(&gameDetails)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &gameDetails, nil
}
