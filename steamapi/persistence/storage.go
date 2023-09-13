package persistence

import (
	"database/sql"
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

type StorageDB interface {
	GetAllFrom(limit int) ([]entity.Item, error)
	Update(item entity.Item) error
	SaveGameDetails(dataProcessed []model.AppDetails) error
	GetGameDetails(id int) (*entity.GameDetails, error)
	GetAllByAppID(appID int) ([]entity.Item, error)
	GetGamesByPage(startIndex, pageSize int) ([]entity.GameDetails, int, error)
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
		fullGameAppID, _ := strconv.Atoi(appDetail.Fullgame.AppID)
		initialPrice, _ := strconv.ParseFloat(strconv.FormatInt(appDetail.PriceOverview.Initial, 10), 64)
		initialFinal, _ := strconv.ParseFloat(strconv.FormatInt(appDetail.PriceOverview.Final, 10), 64)
		query := `
	           INSERT INTO games_details (
	               app_id,
	               name,
	               description,
	               fullgame_app_id,
	               fullgame_name,
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
	               genre_id,
	               type_genre,
	               release_date,
	               coming_soon,
	               currency,
	               initial_price,
	               final_price,
	               discount_percent,
	               formatted_initial_price,
	               formatted_final_price
	           )
	           VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	           ON DUPLICATE KEY UPDATE
	            name = VALUES(name),
	            description = VALUES(description),
	            fullgame_app_id = VALUES(fullgame_app_id),
	            fullgame_name = VALUES(fullgame_name),
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
				genre_id = VALUES(genre_id),
				type_genre = VALUES(type_genre),
	            release_date = VALUES(release_date),
	            coming_soon = VALUES(coming_soon),
	            currency = VALUES(currency),
	           	initial_price = VALUES(initial_price),
	           	final_price = VALUES(final_price),
	        	discount_percent = VALUES(discount_percent),
	            formatted_initial_price = VALUES(formatted_initial_price),
	            formatted_final_price = VALUES(formatted_final_price)
	       `
		_, err := GetDB().Query(
			query,
			appDetail.SteamAppid,
			appDetail.Name,
			appDetail.Description,
			fullGameAppID,
			appDetail.Fullgame.Name,
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
			strings.Join(getGenreIDs(appDetail.Genres), ", "),
			strings.Join(getGenreTypes(appDetail.Genres), ", "),
			appDetail.ReleaseDate.Date,
			appDetail.ReleaseDate.ComingSoon,
			appDetail.PriceOverview.Currency,
			initialPrice,
			initialFinal,
			appDetail.PriceOverview.DiscountPercent,
			appDetail.PriceOverview.InitialFormatted,
			appDetail.PriceOverview.FinalFormatted,
		)
		if err != nil {
			log.Printf("Hubo un error al guardar los juegos: %v\n", err)
			return err
		}
	}

	return nil
}

func getGenreIDs(genres []model.Genre) []string {
	genreIDs := make([]string, len(genres))
	for i, genre := range genres {
		genreIDs[i] = genre.ID
	}
	return genreIDs
}

func getGenreTypes(genres []model.Genre) []string {
	genreTypes := make([]string, len(genres))
	for i, genre := range genres {
		genreTypes[i] = genre.Description
	}
	return genreTypes
}

func mapGenresFromDB(genreID, typeGenre string) []entity.Genre {
	// Dividir los valores por comas para obtener los IDs y descripciones
	idValues := strings.Split(genreID, ", ")
	descriptionValues := strings.Split(typeGenre, ", ")

	// Crear un slice de géneros
	var genres []entity.Genre

	// Asegurarse de que haya la misma cantidad de IDs y descripciones
	if len(idValues) != len(descriptionValues) {
		return genres
	}

	// Mapear los datos a la estructura de géneros
	for i := 0; i < len(idValues); i++ {
		genre := entity.Genre{
			GenreID:   idValues[i],
			TypeGenre: descriptionValues[i],
		}
		genres = append(genres, genre)
	}

	return genres
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
	gameDetails.Genre = mapGenresFromDB(gameDetails.GenreID, gameDetails.TypeGenre)
	return &gameDetails, nil
}

func (s storage) GetGamesByPage(startIndex, pageSize int) ([]entity.GameDetails, int, error) {
	var games []entity.GameDetails

	query := "SELECT * FROM games_details LIMIT ?, ?"
	err := GetDB().Select(&games, query, startIndex, pageSize)
	if err != nil {
		log.Printf("Error al obtener los datos: %v\n", err)
		return nil, 0, err
	}

	totalItems, err := getTotalGamesCount()
	if err != nil {
		log.Printf("Error al obtener el total: %v\n", err)
		return nil, 0, err
	}
	return games, totalItems, nil
}

func getTotalGamesCount() (int, error) {
	var totalItems int
	query := "SELECT COUNT(*) FROM games_details"
	err := GetDB().Get(&totalItems, query)
	if err != nil {
		log.Printf("Hubo un error al obtener el total de los datos: %v\n", err)
		return 0, err
	}
	return totalItems, nil
}
