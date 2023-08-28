package service

//
//import (
//	"context"
//	"encoding/csv"
//	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
//	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
//	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
//	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
//	"log"
//	"os"
//	"strconv"
//	"strings"
//	"sync"
//)
//
//type gameProcessor struct {
//	storage     persistence.Storage
//	steamClient SteamClient
//}
//
//func NewGameProcessor(storage persistence.Storage, steamClient SteamClient) *gameProcessor {
//	return &gameProcessor{
//		storage:     storage,
//		steamClient: steamClient,
//	}
//}
//
//func (sv *gameProcessor) RunProcessData(ctx context.Context, limit int) error {
//	games, err := sv.storage.GetAllFrom(limit)
//	if err != nil {
//		return err
//	}
//
//	gamesDetails, err := sv.getGamesFromAPI(ctx, games)
//	if err != nil {
//		return err
//	}
//
//	dataProcessed, err := sv.processResponse(gamesDetails, games)
//	if err != nil {
//		return err
//	}
//
//	err = sv.saveToCSV(dataProcessed)
//	if err != nil {
//		return err
//	}
//
//	return err
//
//}
//
//func (sv *gameProcessor) getGamesFromAPI(ctx context.Context, items []entity.Item) ([]steamapi.AppDetailsResponse, error) {
//	var wg sync.WaitGroup
//	var processingErrors []error
//	semaphore := make(chan struct{}, 10)
//	processedCount := 0
//	var responseData []steamapi.AppDetailsResponse
//
//	for _, id := range getIds(items) {
//		wg.Add(1)
//		semaphore <- struct{}{}
//		appId := id
//
//		go func(id int64) {
//			defer func() {
//				wg.Done()
//				<-semaphore
//			}()
//
//			response, err := sv.steamClient.GetAppDetails(id)
//			if err != nil {
//				processingErrors = append(processingErrors, err)
//				log.Printf("Error al procesar appID %d: %v\n", id, err)
//				return
//			}
//
//			if response != nil && response.Success {
//				log.Printf("Response for appID %d: %+v\n", id, response)
//				responseData = append(responseData, *response)
//				processedCount++
//			}
//			log.Printf("Elementos procesados hasta ahora: %d", processedCount)
//		}(appId)
//	}
//
//	return responseData, nil
//}
//
//func (sv *gameProcessor) processResponse(responseData []steamapi.AppDetailsResponse, games []entity.Item) ([]steamapi.AppDetails, error) {
//	var appDetails []steamapi.AppDetails
//
//	for _, response := range responseData {
//		data := response.Data
//		appId := data.SteamAppid
//
//		if response.Success {
//			data.SupportedLanguages = utils.ParseSupportedLanguages(data.SupportedLanguagesRaw)
//
//			if data.Type == "game" || data.Type == "dlc" {
//				if err := sv.updateData(games, appId, true); err != nil {
//					log.Printf("Error al actualizar el estado del appID: %v\n", err)
//					return nil, err
//				}
//				appDetails = append(appDetails, data)
//			}
//		} else {
//			if err := sv.updateData(games, appId, false); err != nil {
//				log.Printf("Error al actualizar el estado del appID: %v\n", err)
//				return nil, err
//			}
//		}
//	}
//	return appDetails, nil
//}
//
//func (sv *gameProcessor) updateData(games []entity.Item, id int64, isValid bool) error {
//	findItem := func(games []entity.Item, id int64) entity.Item {
//		for _, game := range games {
//			if game.Appid == id {
//				return game
//			}
//		}
//		return entity.Item{}
//	}
//	itemFound := findItem(games, id)
//	itemFound.IsValid = isValid
//	log.Printf("Insertando juego/appID: %s/%d\n", itemFound.Name, id)
//	return sv.storage.Update(itemFound)
//}
//
//func findItemFrom(games []entity.Item, id int64) entity.Item {
//	for _, game := range games {
//		if game.Appid == id {
//			return game
//		}
//	}
//	return entity.Item{}
//}
//
//func getIds(items []entity.Item) []int64 {
//	var apps []int64
//
//	for _, game := range items {
//		apps = append(apps, game.Appid)
//	}
//	return apps
//
//}
//
//func (sv *gameProcessor) saveToCSV(data []steamapi.AppDetails) error {
//	filePath := "/home/tomi/GCPSteamAnalytics/data/gamesDetails.csv"
//	existingData, err := utils.LoadExistingData(filePath)
//	if err != nil {
//		return err
//	}
//	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	writer := csv.NewWriter(file)
//	defer writer.Flush()
//	// Verificar si el archivo está vacío
//	fileInfo, _ := file.Stat()
//	if fileInfo.Size() == 0 {
//		header := []string{
//			"SteamAppid",
//			"Description",
//			"Type",
//			"Name",
//			"Publishers",
//			"Developers",
//			"isFree",
//			"InterfaceLanguages",
//			"FullAudioLanguages",
//			"SubtitlesLanguages",
//			"Windows",
//			"Mac",
//			"Linux",
//			"Date",
//			"ComingSoon",
//			"Currency",
//			"DiscountPercent",
//			"InitialFormatted",
//			"FinalFormatted",
//		}
//		if err := writer.Write(header); err != nil {
//			return err
//		}
//	}
//	for _, app := range data {
//		if _, exists := existingData[int(app.SteamAppid)]; !exists {
//			record := []string{
//				strconv.Itoa(int(app.SteamAppid)),
//				app.Description,
//				app.Type,
//				app.Name,
//				strings.Join(app.Publishers, ", "),
//				strings.Join(app.Developers, ", "),
//				strconv.FormatBool(app.IsFree),
//				utils.GetSupportedLanguagesString(app.SupportedLanguages["interface"]),
//				utils.GetSupportedLanguagesString(app.SupportedLanguages["full_audio"]),
//				utils.GetSupportedLanguagesString(app.SupportedLanguages["subtitles"]),
//				strconv.FormatBool(app.Platforms.Windows),
//				strconv.FormatBool(app.Platforms.Mac),
//				strconv.FormatBool(app.Platforms.Linux),
//				app.ReleaseDate.Date,
//				strconv.FormatBool(app.ReleaseDate.ComingSoon),
//				app.PriceOverview.Currency,
//				strconv.Itoa(int(app.PriceOverview.DiscountPercent)),
//				utils.FormatInitial(float64(app.PriceOverview.Initial) / 100),
//				app.PriceOverview.FinalFormatted,
//			}
//			if err := writer.Write(record); err != nil {
//				return err
//			}
//			// Agregar el appID al mapa de datos existentes
//			existingData[int(app.SteamAppid)] = true
//		}
//	}
//	return nil
//}
