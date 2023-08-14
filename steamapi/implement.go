package steamapi

// GetAllAppIDs obtiene todos los appid almacenados en la base de datos MySQL.
func (s *SteamAPI) GetAllAppIDs() ([]int, error) {
	query := "SELECT appid FROM games"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appIDs []int
	for rows.Next() {
		var steamAppID int
		if err := rows.Scan(&steamAppID); err != nil {
			return nil, err
		}
		appIDs = append(appIDs, steamAppID)
	}

	return appIDs, nil
}
