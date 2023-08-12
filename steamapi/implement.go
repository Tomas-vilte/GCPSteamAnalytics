package steamapi

//const batchSize = 1000
//
//func (s *SteamAPI) InsertBatchData(items []steamapi.GameDetails) error {
//	if len(items) == 0 {
//		return nil
//	}
//
//	// Creamos la consulta para la inserción en lotes
//	query := "INSERT INTO gamesdetails (steamAppid, nameGame, shortDescription, developers) VALUES "
//	var vals []interface{}
//	for i, item := range items {
//		query += "(?, ?, ?, ?)"
//		// Convertir el slice de developers a una cadena separada por comas
//		developersStr := strings.Join(item.Developers, ",")
//		vals = append(vals, item.SteamAppid, item.NameGame, item.ShortDescription, developersStr)
//		if i < len(items)-1 {
//			query += ", "
//		}
//	}
//
//	// Ejecutamos la consulta en la base de datos
//	_, err := s.DB.Exec(query, vals...)
//	if err != nil {
//		log.Printf("Error al insertar el lote de elementos: %v", err)
//		return err
//	}
//
//	return nil
//}
//
//func (s *SteamAPI) InsertInBatch(items []steamapi.GameDetails) error {
//	// Dividir los datos en lotes
//	numItems := len(items)
//	numBatches := (numItems + batchSize - 1) / batchSize
//
//	// Iterar a través de los lotes y realizar la inserción por lotes
//	for i := 0; i < numBatches; i++ {
//		start := i * batchSize
//		end := (i + 1) * batchSize
//		if end > numItems {
//			end = numItems
//		}
//
//		batchData := items[start:end]
//
//		// Verificar si algún juego en el lote ya existe en la base de datos
//		existingGames := make([]steamapi.GameDetails, 0)
//		for _, item := range batchData {
//			exists, err := s.GameExistsInDatabase(int(item.SteamAppid))
//			if err != nil {
//				return err
//			}
//			if !exists {
//				existingGames = append(existingGames, item)
//			}
//		}
//
//		// Insertar los juegos que no existen en la base de datos
//		err := s.InsertBatchData(existingGames)
//		if err != nil {
//			return err
//		}
//
//		// Registro de logging: Imprimir los juegos insertados en el lote
//		for _, item := range existingGames {
//			log.Printf("Juego insertado en la base de datos: %s (appid: %d)", item.NameGame, item.SteamAppid)
//		}
//	}
//
//	return nil
//}

// GetAppIDs obtiene todos los appid almacenados en la base de datos MySQL.
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

func (s *SteamAPI) GameExistsInDatabase(appid int) (bool, error) {
	query := "SELECT COUNT(*) FROM gamesdetails WHERE steamAppid = ?"
	var count int

	err := s.DB.QueryRow(query, appid).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *SteamAPI) LoadLastProcessedAppid() (int, error) {
	var lastProcessedAppid int
	query := "SELECT last_appid FROM state_table"
	err := s.DB.QueryRow(query).Scan(&lastProcessedAppid)
	if err != nil {
		return 0, err
	}
	return lastProcessedAppid, nil
}

// Función para guardar el último appid procesado en la tabla state_table
func (s *SteamAPI) SaveLastProcessedAppid(lastProcessedAppid int) error {
	query := "UPDATE state_table SET last_appid = ?"
	_, err := s.DB.Exec(query, lastProcessedAppid)
	if err != nil {
		return err
	}
	return nil
}
