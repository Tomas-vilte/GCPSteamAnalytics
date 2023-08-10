package steamapi

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"log"
	"net/http"
	"sync"
	"time"
)

type SteamAPI struct {
	DB *sql.DB
}

const maxAttempts = 10 // Número máximo de intentos de solicitud

func (s *SteamAPI) ExtractAndSaveLimitedGameDetails(limit int) error {
	// Cargar el último appid procesado desde la base de datos
	lastProcessedAppID, err := s.LoadLastProcessedAppid()
	if err != nil {
		log.Printf("Error al cargar el último appID procesado: %v", err)
		return err
	}
	lastSuccessfulAppID := lastProcessedAppID
	var processedCount int
	// Obtener los appids desde la base de datos a partir del último procesado
	appids, err := s.GetAppIDs(lastProcessedAppID)
	if err != nil {
		log.Printf("Error al obtener los AppIDs: %v", err)
		return err
	}

	// Utilizar un semáforo para controlar la concurrencia y el número máximo de solicitudes simultáneas
	semaphore := make(chan struct{}, 10) // Permite 10 solicitudes simultáneas

	var wg sync.WaitGroup

	client := http.Client{}
	var count int
	var missingDataCount int

	// Crear un canal para recibir los detalles de los juegos
	gameDetailsChan := make(chan steamapi.GameDetails, 10) // Tamaño del canal debe ser mayor que el tamaño de lote

	// Crear una goroutine para insertar los detalles de los juegos en lotes
	go func() {
		var batchData []steamapi.GameDetails

		for gameDetails := range gameDetailsChan {
			batchData = append(batchData, gameDetails)
			if len(batchData) >= 10 { // Tamaño del lote
				if err := s.InsertInBatch(batchData); err != nil {
					log.Printf("Error al insertar lote de datos: %v\n", err)
				}
				batchData = nil
			}
		}

		// Insertar los datos restantes en el último lote
		if len(batchData) > 0 {
			if err := s.InsertInBatch(batchData); err != nil {
				log.Printf("Error al insertar lote de datos final: %v\n", err)
			}
		}
	}()

	for _, appid := range appids {
		if count >= limit {
			break // Salir del bucle una vez alcanzado el límite
		}
		wg.Add(1)
		semaphore <- struct{}{} // Bloquea si ya hay 10 solicitudes simultáneas
		go func(appid int) {
			defer func() {
				<-semaphore // Libera un slot del semáforo para permitir otra solicitud
				wg.Done()
			}()

			count++
			url := fmt.Sprintf("https://store.steampowered.com/api/appdetails?l=spanish&appids=%d", appid)

			for attempt := 0; attempt < maxAttempts; attempt++ {
				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()
				// Realiza la solicitud HTTP con el cliente configurado
				req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
				resp, err := client.Do(req)
				if err != nil {
					log.Printf("Error al obtener detalles para appid %d (intentando de nuevo): %v\n", appid, err)
					continue
				}

				if resp.StatusCode == http.StatusTooManyRequests {
					log.Printf("Obtenido código de estado 429 (demasiadas solicitudes). Esperando 60 segundos...\n")
					time.Sleep(60 * time.Second)
					continue
				}
				defer resp.Body.Close()

				time.Sleep(10 * time.Second)

				var steamResponse map[string]steamapi.SteamApiResponse
				if err := json.NewDecoder(resp.Body).Decode(&steamResponse); err != nil {
					log.Printf("Error al decodificar la respuesta %d: %v\n", appid, err)
					return
				}

				appidResponse := steamResponse[fmt.Sprintf("%d", appid)]

				// Verifica si existe la key data
				if appidResponse.Success && appidResponse.Data != nil {
					lastSuccessfulAppID = appid
					var gameDetails steamapi.GameDetails
					if err := json.Unmarshal(appidResponse.Data, &gameDetails); err != nil {
						log.Printf("Error al analizar los detalles del juego para appid %d: %v\n", appid, err)
						return
					}
					processedCount++
					// Agrega el gameDetails al slice de juegos a insertar
					gameDetailsChan <- gameDetails
					log.Printf("Detalles de juegos insertados correctamente. Juegos procesados en total: %d\n", processedCount)

				} else {
					missingDataCount++
					log.Printf("Error al obtener detalles para appid %d: el success es false\n", appid)
				}
				break
			}
		}(appid)
	}

	wg.Wait()
	close(semaphore)
	close(gameDetailsChan)

	// Actualiza la tabla state_table con el último appID exitoso, incluso si hay errores
	if lastSuccessfulAppID > lastProcessedAppID {
		if err := s.SaveLastProcessedAppid(lastSuccessfulAppID); err != nil {
			log.Printf("Error al guardar el último appID procesado: %v", err)
		}
	}
	log.Printf("Total de detalles sin la clave \"data\": %d\n", missingDataCount)
	return nil
}
