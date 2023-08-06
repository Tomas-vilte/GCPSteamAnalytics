package steamapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tomas-vilte/GCPSteamAnalytics/db"
	"github.com/Tomas-vilte/GCPSteamAnalytics/models"
)

type GameDetails struct {
	models.StoreItem
}

type SteamAPI struct {
	Dba db.Database
}

type Response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type SteamData interface {
	ExtractGameDetails() ([]GameDetails, error)
}

func (s *SteamAPI) ExtractGameDetails() ([]GameDetails, error) {
	// Obtener los appids y nombres desde la base de datos
	items, err := s.Dba.GetAppIDs()
	if err != nil {
		return nil, err
	}

	// Realizar una solicitud HTTP para obtener los detalles de los juegos para cada appid
	var gamesDetails []GameDetails
	requestCount := 0
	for _, item := range items {
		if requestCount >= 10 {
			break
		}
		appid := item
		url := fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%d", appid)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error al obtener detalles para appid %d: %v\n\n", appid, err)
			continue
		}
		defer resp.Body.Close()

		var steamResponse map[string]Response
		if err := json.NewDecoder(resp.Body).Decode(&steamResponse); err != nil {
			log.Printf("Error al decodificar la respuesta %d: %v\n\n", appid, err)
			continue
		}

		appidResponse := steamResponse[fmt.Sprintf("%d", appid)]

		// Verifica si existe la key data
		if appidResponse.Success {
			var gameDetails GameDetails
			if err := json.Unmarshal(appidResponse.Data, &gameDetails); err != nil {
				log.Printf("Error al analizar los detalles del juego para appid %d: %v\n\n", appid, err)
				continue
			}

			log.Printf("ID de cada juego %d:\n\n\n", appid)
			log.Printf("Nombre de cada juego: %s\n\n", gameDetails.NameGame)
			log.Printf("Descripcion de cada juego: %s\n\n", gameDetails.ShortDescription)

			// Agrega los detalles del juego al slice 'gamesDetails'
			gamesDetails = append(gamesDetails, gameDetails)
		} else {
			log.Printf("Error al obtener detalles para appid %d: el success es false\n\n", appid)
		}
		requestCount++
	}

	return gamesDetails, nil
}
