package steamapi

import (
	"strings"
	"sync"
)

// GetAllAppIDs obtiene todos los appIDs almacenados en la base de datos MySQL
// que son mayores que el último appID procesado.
func (s *SteamAPI) GetAllAppIDs(lastProcessedAppID int) ([]int, error) {
	query := "SELECT appid FROM games WHERE appid > ?"
	rows, err := s.DB.Query(query, lastProcessedAppID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appids []int
	for rows.Next() {
		var steamAppid int
		if err := rows.Scan(&steamAppid); err != nil {
			return nil, err
		}
		appids = append(appids, steamAppid)
	}

	return appids, nil
}

// LoadLastProcessedAppid carga el último appID procesado desde la tabla state_table.
func (s *SteamAPI) LoadLastProcessedAppid() (int, error) {
	var lastProcessedAppid int
	query := "SELECT last_appid FROM state_table"
	err := s.DB.QueryRow(query).Scan(&lastProcessedAppid)
	if err != nil {
		return 0, err
	}
	return lastProcessedAppid, nil
}

// SaveLastProcessedAppid guarda el último appID procesado en la tabla state_table.
func (s *SteamAPI) SaveLastProcessedAppid(lastProcessedAppid int) error {
	query := "UPDATE state_table SET last_appid = ?"
	_, err := s.DB.Exec(query, lastProcessedAppid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SteamAPI) AreEmptyAppIDs(appIDs []int) (map[int]bool, error) {
	query := "SELECT appid FROM empty_appids WHERE appid IN (?)"

	batchSize := 300
	emptyAppIDs := new(sync.Map)

	for i := 0; i < len(appIDs); i += batchSize {
		end := i + batchSize
		if end > len(appIDs) {
			end = len(appIDs)
		}

		batchIDs := appIDs[i:end]
		placeholders := make([]string, len(batchIDs))
		args := make([]interface{}, len(batchIDs))

		for j, id := range batchIDs {
			placeholders[j] = "?"
			args[j] = id
		}

		batchQuery := strings.Replace(query, "(?)", "("+strings.Join(placeholders, ",")+")", 1)
		rows, err := s.DB.Query(batchQuery, args...)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var appID int
			err := rows.Scan(&appID)
			if err != nil {
				return nil, err
			}
			emptyAppIDs.Store(appID, true)
		}

		rows.Close() // Cierra las filas en cada iteración
	}

	resultMap := make(map[int]bool)
	emptyAppIDs.Range(func(key, value interface{}) bool {
		resultMap[key.(int)] = value.(bool)
		return true
	})

	return resultMap, nil
}

func (s *SteamAPI) AddToEmptyAppIDsTable(appID int) error {
	query := "INSERT INTO empty_appids (appid) VALUES (?)"
	_, err := s.DB.Exec(query, appID)
	return err
}

func (s *SteamAPI) GetStartIndexToProcess(lastProcessedAppID int, appIDs []int) int {
	startIndex := 0
	if lastProcessedAppID != 0 {
		for i, appID := range appIDs {
			if appID == lastProcessedAppID {
				startIndex = i + 1
				break
			}
		}
	}
	return startIndex
}
