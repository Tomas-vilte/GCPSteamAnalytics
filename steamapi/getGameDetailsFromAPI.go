package steamapi

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	baseURL   = "https://store.steampowered.com/api/appdetails"
	apiKey    = "YOUR_STEAM_API_KEY"
	language  = "spanish"
	outputCSV = "steam_data.csv"
	cc        = "AR"
)

type SteamAPI struct {
	DB *sql.DB
}

type AppDetailsResponse struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

type AppDetails struct {
	SteamAppid  int64    `json:"steam_appid"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"short_description"`
	Developers  []string `json:"developers"`
	Publishers  []string `json:"publishers"`
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	PriceOverview struct {
		Currency         string `json:"currency"`
		DiscountPercent  int    `json:"discount_percent"`
		InitialFormatted string `json:"initial_formatted"`
		FinalFormatted   string `json:"final_formatted"`
	} `json:"price_overview"`
}

func getSteamData(appIDs []int) ([]AppDetails, error) {
	var wg sync.WaitGroup
	var results []AppDetails
	var errors []error

	// Crear un semáforo con un límite de 10 solicitudes concurrentes
	semaphore := make(chan struct{}, 10)

	for i, appID := range appIDs {
		wg.Add(1)
		semaphore <- struct{}{} // Adquirir un lugar en el semáforo
		go func(id int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Liberar un lugar en el semáforo

			url := fmt.Sprintf("%s?l=%s&appids=%d&key=%s&cc=%s", baseURL, language, id, apiKey, cc)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				errors = append(errors, err)
				return
			}

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				errors = append(errors, err)
				return
			}
			defer response.Body.Close()

			var responseData map[string]AppDetailsResponse
			err = json.NewDecoder(response.Body).Decode(&responseData)
			if err != nil {
				errors = append(errors, err)
				return
			}

			if responseData[strconv.Itoa(id)].Success {
				result := responseData[strconv.Itoa(id)].Data
				if result.Type == "game" || result.Type == "dlc" {
					results = append(results, result)
					log.Printf("Insertado juego/appID: %s/%d\n", result.Name, id)
				} else {
					log.Printf("No insertado (tipo no válido)/appID: %d\n", id)
				}
			}
		}(appID)

		// Aplicar un timeout después de cada grupo de 10 solicitudes
		if (i+1)%10 == 0 || i == len(appIDs)-1 {
			wg.Wait()                    // Esperar a que las solicitudes en progreso terminen antes de continuar
			time.Sleep(10 * time.Second) // Aplicar el timeout de 10 segundos
		}
	}

	if len(errors) > 0 {
		return nil, errors[0]
	}

	return results, nil
}

func saveToCSV(data []AppDetails, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"SteamAppid", "Description", "Type", "Name", "Publishers", "Developers", "Windows", "Mac",
		"Linux", "date", "comingSoon",
		"currency", "discount_percent", "initial_formatted", "final_formatted"})

	for _, entry := range data {
		publishers := strings.Join(entry.Publishers, ", ")
		developers := strings.Join(entry.Developers, ", ")
		writer.Write([]string{
			strconv.FormatInt(entry.SteamAppid, 10),
			entry.Description,
			entry.Type,
			entry.Name,
			publishers,
			developers,
			strconv.FormatBool(entry.Platforms.Windows),
			strconv.FormatBool(entry.Platforms.Mac),
			strconv.FormatBool(entry.Platforms.Linux),
			entry.ReleaseDate.Date,
			strconv.FormatBool(entry.ReleaseDate.ComingSoon),
			entry.PriceOverview.Currency,
			strconv.Itoa(entry.PriceOverview.DiscountPercent),
			entry.PriceOverview.InitialFormatted,
			entry.PriceOverview.FinalFormatted,
		})
	}

	return nil
}

func (s *SteamAPI) RunSteamDataExtraction(appids []int) error {
	data, err := getSteamData(appids)
	if err != nil {
		return fmt.Errorf("error getting Steam data: %v", err)
	}

	err = saveToCSV(data, outputCSV)
	if err != nil {
		return fmt.Errorf("error saving data to CSV: %v", err)
	}

	return nil
}
