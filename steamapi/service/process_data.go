package service

import (
	"context"
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
	"log"
	"net/http"
	"sync"
)

type GameProcessor struct {
	storage  persistence.Storage
	steamAPI steamapi.SteamData
}

func (sv *GameProcessor) RunProcessData(ctx context.Context, limit int) error {
	games, err := sv.storage.GetAllFrom(limit)
	if err != nil {
		return err
	}

	gamesDetails, err := sv.getGamesFromAPI(ctx, games)

	return err

}

func (sv *GameProcessor) getGamesFromAPI(ctx context.Context, items []handlers.Item) ([]steamapi.AppDetails, error) {
	var wg sync.WaitGroup
	var processedData []steamapi.AppDetails
	var processingErrors []error
	semaphore := make(chan struct{}, 10)
	processedCount := 0

	for _, id := range getIds(items) {
		wg.Add(1)
		semaphore <- struct{}{}
		appId := id

		go func(id int64) {
			defer func() {
				wg.Done()
				<-semaphore
			}()

			sv.steamAPI.ProcessAppID(id)
		}(appId)
	}

	return nil, nil
}

func getIds(items []handlers.Item) []int64 {
	var apps []int64

	for _, game := range items {
		apps = append(apps, game.Appid)
	}
	return apps

}

func (sv *GameProcessor) UNAME(id int64, isValid bool) (*steamapi.AppDetailsResponse, error) {
	appDetailsResponse, err := sv.steamAPI.ProcessAppID(id)
	if err != nil {
		// Manejar el error de la llamada API
		return nil, err
	}
	if appDetailsResponse. {
		// La respuesta de la API fue exitosa, puedes acceder a los datos
		data := appDetailsResponse.Data
		data.SupportedLanguages = utils.ParseSupportedLanguages(data.SupportedLanguagesRaw)

		if data.Type == "game" || data.Type == "dlc" {
			log.Printf("Insertando juego/appID: %s/%d\n", data.Name, id)
			err = sv.storage.Update(handlers.Item{Appid: id, IsValid: true})
			if err != nil {
				log.Printf("Error al actualizar el estado del appID: %v\n", err)
			}
			return &data, nil
		} else {
			err := sv.storage.Update(handlers.Item{Appid: id, IsValid: false})
			if err != nil {
				log.Printf("Error al actualizar el estado del appID: %v\n", err)
			}
			log.Printf("No insertado (tipo no válido:%s) / appID: %d\n", data.Type, id)
		}
	} else {
		// La respuesta de la API no fue exitosa, maneja el caso aquí
		log.Printf("Llamada a API no exitosa para appID: %d\n", id)
	}

	return nil, nil
}
