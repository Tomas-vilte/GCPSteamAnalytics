package steamapi

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
