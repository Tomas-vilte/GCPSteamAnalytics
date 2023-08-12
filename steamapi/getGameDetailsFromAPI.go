package steamapi

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	baseURL  = "https://store.steampowered.com/api/appdetails"
	apiKey   = "1A059D89640D054BB20FF254FB529E14"
	language = "spanish"
	cc       = "AR"
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
		DiscountPercent  int64  `json:"discount_percent"`
		InitialFormatted string `json:"initial_formatted"`
		FinalFormatted   string `json:"final_formatted"`
	} `json:"price_overview"`
}

func GetSteamData(appIDs []int, limit int) ([]AppDetails, error) {
	var wg sync.WaitGroup
	var results []AppDetails
	var errors []error

	// Crear un semáforo con un límite de 10 solicitudes concurrentes
	semaphore := make(chan struct{}, 10)
	for i, appID := range appIDs {
		if len(results) >= limit {
			break
		}
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
					log.Printf("No insertado (tipo no válido: %s)/appID: %d\n", result.Type, id)
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

func SaveToCSV(data []AppDetails, filePath string) error {
	existingData, err := loadExistingData(filePath)
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
				strconv.FormatBool(app.Platforms.Windows),
				strconv.FormatBool(app.Platforms.Mac),
				strconv.FormatBool(app.Platforms.Linux),
				app.ReleaseDate.Date,
				strconv.FormatBool(app.ReleaseDate.ComingSoon),
				app.PriceOverview.Currency,
				strconv.Itoa(int(app.PriceOverview.DiscountPercent)),
				app.PriceOverview.InitialFormatted,
				app.PriceOverview.FinalFormatted,
				// ... otros campos que quieras guardar
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

func loadExistingData(filePath string) (map[int]bool, error) {
	existingData := make(map[int]bool)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return existingData, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Leer y descartar la primera fila (encabezados)
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Leer las filas restantes y procesar los appIDs
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Asegurarse de que haya al menos un valor en el registro antes de convertir
		if len(record) < 1 {
			continue
		}

		appID, err := strconv.Atoi(record[0])
		if err != nil {
			// Puede ser útil agregar un registro de depuración aquí para identificar registros incorrectos
			continue // Saltar esta fila y seguir con la siguiente
		}
		existingData[appID] = true
	}

	return existingData, nil
}
