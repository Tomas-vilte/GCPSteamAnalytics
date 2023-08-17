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

func (s *SteamAPI) IsEmptyAppID(appID int) (bool, error) {
	query := "SELECT COUNT(*) FROM empty_appids WHERE appid = ?"
	var count int
	err := s.DB.QueryRow(query, appID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
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
