package persistence

import (
	"database/sql"
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
	GetGameDetails(id int) (*GameDetails, error)
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

type GameDetails struct {
	AppID              int    `json:"app_id"`
	Description        string `json:"description"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	Publishers         string `json:"publishers"`
	Developers         string `json:"developers"`
	IsFree             bool   `json:"is_free"`
	InterfaceLanguages string `json:"interface_languages"`
	FullAudioLanguages string `json:"full_audio_languages"`
	SubtitlesLanguages string `json:"subtitles_languages"`
	Windows            bool   `json:"windows"`
	Mac                bool   `json:"mac"`
	Linux              bool   `json:"linux"`
	ReleaseDate        struct {
		Date       string `json:"date"`
		ComingSoon bool   `json:"coming_soon"`
	} `json:"release_date"`
	Currency         string `json:"currency"`
	DiscountPercent  int    `json:"discount_percent"`
	InitialFormatted string `json:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted"`
}

func (s storage) GetGameDetails(gameID int) (*GameDetails, error) {
	// Consulta SQL para obtener los detalles del juego por su ID.
	query := `
       SELECT
           app_id,
           description,
           type,
           name,
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
       FROM
           games_details
       WHERE
           app_id = ?
   `

	// Ejecutar la consulta SQL y escanear los resultados en una estructura AppDetails.
	var gameDetails GameDetails
	err := GetDB().QueryRow(query, gameID).Scan(
		&gameDetails.AppID,
		&gameDetails.Description,
		&gameDetails.Type,
		&gameDetails.Name,
		&gameDetails.Publishers,
		&gameDetails.Developers,
		&gameDetails.IsFree,
		&gameDetails.InterfaceLanguages,
		&gameDetails.FullAudioLanguages,
		&gameDetails.SubtitlesLanguages,
		&gameDetails.Windows,
		&gameDetails.Mac,
		&gameDetails.Linux,
		&gameDetails.ReleaseDate.Date,
		&gameDetails.ReleaseDate.ComingSoon,
		&gameDetails.Currency,
		&gameDetails.DiscountPercent,
		&gameDetails.InitialFormatted,
		&gameDetails.FinalFormatted,
	)

	if err != nil {
		log.Printf("error: %v", err)
		if err == sql.ErrNoRows {
			log.Printf("error: %v", err)
			return nil, nil // No se encontraron detalles del juego en la base de datos.
		}
		log.Printf("error: %v", err)
		return nil, err // Ocurrió un error diferente al ejecutar la consulta SQL.
	}

	return &gameDetails, nil
}
