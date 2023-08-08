package steamapi

import (
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"log"
)

const batchSize = 100

func (s *SteamAPI) InsertBatch(items []steamapi.GameDetails) error {

	// Dividimos los datos en lotes
	numItems := len(items)
	numBatches := (numItems + batchSize - 1) / batchSize

	// Iteramos los lotes y realizamos la inserción por lotes
	for i := 0; i < numBatches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize
		if end > numItems {
			end = numItems
		}

		batchData := items[start:end]

		err := s.InsertBatchData(batchData)
		if err != nil {
			log.Printf("Error al insertar el lote de elementos: %v", err)
			return err
		}
	}

	return nil
}

func (s *SteamAPI) InsertBatchData(items []steamapi.GameDetails) error {
	if len(items) == 0 {
		return nil
	}

	// Creamos la consulta para la insercion en lotes
	query := "INSERT INTO gamesdetails (steamAppid, nameGame, shortDescription) VALUES "
	var vals []interface{}
	for i, item := range items {
		query += "(?, ?, ?)"
		vals = append(vals, item.SteamAppid, item.NameGame, item.ShortDescription)
		if i < len(items)-1 {
			query += ", "
		}
	}

	// Ejecutamos la consulta en la base de datos
	_, err := s.DB.Exec(query, vals...)
	if err != nil {
		log.Printf("Error al insertar el lote de elementos: %v", err)
		return err
	}

	return nil
}

func (s *SteamAPI) InsertInBatch(items []steamapi.GameDetails) error {
	for _, item := range items {
		// Verificar si el juego ya existe en la base de datos
		exists, err := s.GameExistsInDatabase(int(item.SteamAppid))
		if err != nil {
			return err
		}

		// Si el juego ya existe, continuar con el siguiente
		if exists {
			continue
		}

		// Insertar el juego en la base de datos
		err = s.InsertBatch([]steamapi.GameDetails{item})
		if err != nil {
			return err
		}

		// Registro de logging: Imprimir el juego insertado
		log.Printf("Juego insertado en la base de datos: %s (appid: %d)", item.NameGame, item.SteamAppid)
	}

	return nil
}

// GetAppIDs obtiene todos los appid almacenados en la base de datos MySQL.
func (s *SteamAPI) GetAppIDs() ([]int, error) {
	rows, err := s.DB.Query("SELECT appid FROM games")
	if err != nil {
		log.Printf("Error al obtener los appid desde la base de datos: %v", err)
		return nil, err
	}
	defer rows.Close()

	var appids []int
	for rows.Next() {
		var appid int
		if err := rows.Scan(&appid); err != nil {
			log.Printf("Error al escanear el appid desde la base de datos: %v", err)
			return nil, err
		}
		appids = append(appids, appid)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al obtener los appid desde la base de datos: %v", err)
		return nil, err
	}

	return appids, nil
}

func (s *SteamAPI) GameExistsInDatabase(appid int) (bool, error) {
	query := "SELECT COUNT(*) FROM gamesdetails WHERE steamAppid = ?"
	var count int

	err := s.DB.QueryRow(query, appid).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
