package steamapi

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	"time"
)

// GetAllAppIDs obtiene todos los appIDs almacenados en la base de datos MySQL
// que son mayores que el último appID procesado.
func (s *SteamAPI) GetAllAppIDs(limit int) ([]int, error) {
	query := "SELECT app_id FROM game WHERE status = 'PENDING' AND valid = false ORDER BY id LIMIT ?"
	rows, err := s.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appIDs []int
	for rows.Next() {
		var appID int
		err := rows.Scan(&appID)
		if err != nil {
			return nil, err
		}
		appIDs = append(appIDs, appID)
	}

	return appIDs, nil
}

func (s *SteamAPI) UpdateAppStatus(id int, isValid bool) error {
	query := "UPDATE game SET status = ?, valid = ?, updated_at = ? WHERE app_id = ?"
	_, err := s.DB.Exec(query, handlers.PROCESSED, isValid, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
