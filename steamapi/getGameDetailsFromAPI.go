package steamapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"log"
	"net/http"
	"sync"
)

type SteamAPI struct {
	DB *sql.DB
}

func (s *SteamAPI) ExtractAndSaveLimitedGameDetails(limit int) error {
	// Obtener los appids desde la base de datos
	appids, err := s.GetAppIDs()
	if err != nil {
		return err
	}

	// Utilizar un semáforo para controlar la concurrencia y el número máximo de solicitudes simultáneas
	semaphore := make(chan struct{}, 2) // Permite 10 solicitudes simultáneas

	var wg sync.WaitGroup

	var gamesDetails []steamapi.GameDetails // Slice para almacenar los datos a insertar en la base de datos

	// Crear un canal para controlar el cierre del semáforo
	done := make(chan struct{})
	defer close(done) // Cerrar el semáforo al finalizar

	var count int
	for _, appid := range appids {
		if count >= limit {
			break // Salir del bucle una vez alcanzado el límite
		}
		wg.Add(1)
		semaphore <- struct{}{} // Bloquea si ya hay 10 solicitudes simultáneas

		go func(appid int64) {
			defer func() {
				<-semaphore // Libera un slot del semáforo para permitir otra solicitud
				wg.Done()
			}()

			count++

			url := fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%d", appid)
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error al obtener detalles para appid %d: %v\n", appid, err)
				return
			}
			defer resp.Body.Close()

			var steamResponse map[string]steamapi.SteamApiResponse
			if err := json.NewDecoder(resp.Body).Decode(&steamResponse); err != nil {
				log.Printf("Error al decodificar la respuesta %d: %v\n", appid, err)
				return
			}

			appidResponse := steamResponse[fmt.Sprintf("%d", appid)]

			// Verifica si existe la key data
			if appidResponse.Success && appidResponse.Data != nil {
				var gameDetails steamapi.GameDetails
				if err := json.Unmarshal(appidResponse.Data, &gameDetails); err != nil {
					log.Printf("Error al analizar los detalles del juego para appid %d: %v\n", appid, err)
					return
				}

				// Agregar el gameDetails al slice de juegos a insertar
				gamesDetails = append(gamesDetails, gameDetails)
				fmt.Println(gameDetails.SteamAppid, gameDetails.NameGame)

			} else {
				log.Printf("Error al obtener detalles para appid %d: el success es false\n", appid)
			}
		}(int64(appid))
	}

	wg.Wait()

	close(semaphore)

	//// Insertar los datos en la base de datos utilizando el método InsertBatch con goroutines
	if err := s.InsertInBatch(gamesDetails); err != nil {
		return fmt.Errorf("error al guardar los detalles de los juegos en la base de datos: %v", err)
	}
	return nil
}
