package service

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
)

type gameProcessor struct {
	storage     persistence.StorageDB
	steamClient SteamClient
}

func NewGameProcessor(storage persistence.StorageDB, steamClient SteamClient) *gameProcessor {
	return &gameProcessor{
		storage:     storage,
		steamClient: steamClient,
	}
}

func (sv *gameProcessor) RunProcessData(ctx context.Context, limit int) error {
	games, err := sv.storage.GetAllFrom(limit)
	if err != nil {
		return err
	}

	gamesDetails, err := sv.GetGamesFromAPI(ctx, games)
	if err != nil {
		return err
	}

	dataProcessed, err := sv.processResponse(gamesDetails, games)
	if err != nil {
		return err
	}

	err = sv.saveToCSV(dataProcessed)
	if err != nil {
		return err
	}

	return err

}

func (sv *gameProcessor) GetGamesFromAPI(ctx context.Context, items []entity.Item) ([][]byte, error) {
	var wg sync.WaitGroup
	var processingErrors []error
	semaphore := make(chan struct{}, 10)
	var responseData [][]byte

	for _, appId := range getIds(items) {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(id int64) {
			defer func() {
				wg.Done()
				<-semaphore
			}()

			select {
			case <-ctx.Done():
				return
			default:
				response, err := sv.steamClient.GetAppDetails(int(id))
				if err != nil {
					processingErrors = append(processingErrors, err)
					log.Printf("Error al procesar appID %d: %v\n", id, err)
					return
				}

				if response != nil {
					responseData = append(responseData, response)
				}
			}
		}(appId)
	}

	if len(processingErrors) > 0 {
		return nil, fmt.Errorf("ocurrio un error durante los api calls")
	}

	wg.Wait()
	return responseData, nil
}

func (sv *gameProcessor) processResponse(responseData [][]byte, games []entity.Item) ([]steamapi.AppDetails, error) {
	var appDetails []steamapi.AppDetails
	logCounter := 1

	for _, responseBytes := range responseData {
		var responseMap map[string]steamapi.AppDetailsResponse
		if err := json.Unmarshal(responseBytes, &responseMap); err != nil {
			log.Printf("Error al decodificar JSON de la respuesta: %v\n", err)
			continue
		}

		for appIDStr, response := range responseMap {
			data := response.Data
			appID, err := strconv.Atoi(appIDStr)
			if err != nil {
				log.Printf("Error al convertir appID a entero: %v\n", err)
				continue
			}

			if response.Success && (data.Type == "game" || data.Type == "dlc") {
				log.Printf("[%d] Insertando juego/appID: %s/%d\n", logCounter, data.Name, appID)
				appDetails = append(appDetails, data)
			} else {
				log.Printf("[%d] No insertado (tipo no válido: %s) / appID: %d\n", logCounter, data.Type, appID)
			}

			err = sv.updateData(games, int64(appID), response.Success)
			if err != nil {
				log.Printf("[%d] Error al actualizar el estado del appID: %v\n", logCounter, err)
				return nil, err
			}

			log.Printf("[%d] Estado actualizado del juego con appID %d\n", logCounter, appID)

			logCounter++
		}
	}

	return appDetails, nil
}

func (sv *gameProcessor) updateData(games []entity.Item, id int64, isValid bool) error {
	findItem := func(games []entity.Item, id int64) *entity.Item {
		for i := range games {
			if games[i].Appid == id {
				return &games[i]
			}
		}
		return nil
	}

	itemFound := findItem(games, id)
	if itemFound == nil {
		return fmt.Errorf("juego con el appID %d no se encuentra", id)
	}

	itemFound.IsValid = isValid

	if err := sv.storage.Update(*itemFound); err != nil {
		log.Printf("Error al actualizar el estado del juego con appID %d: %v\n", itemFound.Appid, err)
		return err
	}

	return nil
}

func getIds(items []entity.Item) []int64 {
	var apps []int64

	for _, game := range items {
		apps = append(apps, game.Appid)
	}
	return apps

}

func (sv *gameProcessor) saveToCSV(data []steamapi.AppDetails) error {
	filePath := "/home/tomi/GCPSteamAnalytics/data/gamesDetails.csv"
	existingData, err := utils.LoadExistingData(filePath)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Verificar si el archivo está vacío
	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		header := []string{
			"SteamAppid",
			"Description",
			"Type",
			"Name",
			"Publishers",
			"Developers",
			"isFree",
			"InterfaceLanguages",
			"FullAudioLanguages",
			"SubtitlesLanguages",
			"Windows",
			"Mac",
			"Linux",
			"Date",
			"ComingSoon",
			"Currency",
			"DiscountPercent",
			"InitialFormatted",
			"FinalFormatted",
		}
		if err := writer.Write(header); err != nil {
			return err
		}
	}
	for _, app := range data {
		if _, exists := existingData[int(app.SteamAppid)]; !exists {
			record := []string{
				strconv.Itoa(int(app.SteamAppid)),
				app.Description,
				app.Type,
				app.Name,
				strings.Join(app.Publishers, ", "),
				strings.Join(app.Developers, ", "),
				strconv.FormatBool(app.IsFree),
				utils.GetSupportedLanguagesString(app.SupportedLanguages["interface"]),
				utils.GetSupportedLanguagesString(app.SupportedLanguages["full_audio"]),
				utils.GetSupportedLanguagesString(app.SupportedLanguages["subtitles"]),
				strconv.FormatBool(app.Platforms.Windows),
				strconv.FormatBool(app.Platforms.Mac),
				strconv.FormatBool(app.Platforms.Linux),
				app.ReleaseDate.Date,
				strconv.FormatBool(app.ReleaseDate.ComingSoon),
				app.PriceOverview.Currency,
				strconv.Itoa(int(app.PriceOverview.DiscountPercent)),
				utils.FormatInitial(float64(app.PriceOverview.Initial) / 100),
				app.PriceOverview.FinalFormatted,
			}
			if err := writer.Write(record); err != nil {
				return err
			}
			// Agregar el appID al mapa de datos existentes
			existingData[int(app.SteamAppid)] = true
		}
	}
	return nil
}
