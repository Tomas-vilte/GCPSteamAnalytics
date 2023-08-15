package steamapi

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SteamAPI es una estructura que maneja la comunicación con la API de Steam y la base de datos.
type SteamAPI struct {
	DB     *sql.DB
	Client HTTPClient
}

// RunProcessData es la función principal que coordina el proceso de obtención, procesamiento y guardado de datos de Steam.
// Acepta una interfaz 'api' que debe cumplir con la interfaz 'SteamData' definida en models.go.
// 'limit' es la cantidad máxima de juegos a procesar.
func RunProcessData(api steamapi.SteamData, limit int) error {
	// Cargar el último appID procesado.
	lastProcessedAppID, err := api.LoadLastProcessedAppid()
	if err != nil {
		return err
	}

	// Cargar SteamAppIDs previamente procesados
	appIDs, err := api.GetAllAppIDs(lastProcessedAppID)
	if err != nil {
		return err
	}

	// Procesar datos de Steam y obtener los detalles de los juegos.
	data, err := api.ProcessSteamData(appIDs, limit)
	if err != nil {
		return err
	}

	// Guardar los datos procesados en un archivo CSV.
	err = api.SaveToCSV(data, "../data/dataDetails.csv")
	if err != nil {
		return err
	}

	return nil
}

// ProcessSteamData procesa los datos de los juegos de Steam.
// 'appIDs' es una lista de appIDs a procesar, 'limit' es la cantidad máxima de juegos a procesar.
// Retorna los detalles de los juegos procesados y posibles errores de procesamiento.
func (s *SteamAPI) ProcessSteamData(appIDs []int, limit int) ([]steamapi.AppDetails, error) {
	var wg sync.WaitGroup
	var processedData []steamapi.AppDetails
	var processingErrors []error
	semaphore := make(chan struct{}, 10)

	for i, appID := range appIDs {
		if len(processedData) >= limit {
			break
		}

		wg.Add(1)
		semaphore <- struct{}{} // Adquirir un lugar en el semáforo

		go func(id int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Liberar un lugar en el semáforo

			data, err := s.ProcessAppID(id)
			if err != nil {
				processingErrors = append(processingErrors, err)
				return
			}

			if data != nil {
				processedData = append(processedData, *data)
			}
		}(appID)

		if (i+1)%10 == 0 || i == len(appIDs)-1 {
			wg.Wait()
			time.Sleep(10 * time.Second)
		}
	}

	if len(processingErrors) > 0 {
		return nil, processingErrors[0]
	}

	return processedData, nil
}

// ProcessAppID procesa un appID específico y devuelve sus detalles si es un juego válido.
// 'id' es el appID a procesar.
// Retorna los detalles del juego y un posible error si ocurre.
func (s *SteamAPI) ProcessAppID(id int) (*steamapi.AppDetails, error) {
	url := fmt.Sprintf("%s?l=%s&appids=%d&key=%s&cc=%s", baseURL, language, id, apiKey, cc)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseData map[string]steamapi.AppDetailsResponse
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	if responseData[strconv.Itoa(id)].Success {
		data := responseData[strconv.Itoa(id)].Data
		if data.Type == "game" || data.Type == "dlc" {
			log.Printf("Insertando juego/appID: %s/%d\n", data.Name, id)
			err = s.SaveLastProcessedAppid(id)
			if err != nil {
				log.Printf("Error al guardar el último appid procesado: %v\n", err)
			}
			return &data, nil
		} else {
			log.Printf("No insertado (tipo no válido:%s) / appID: %d\n", data.Type, id)
		}
	}

	return nil, nil
}

// SaveToCSV guarda los detalles de los juegos en un archivo CSV.
// 'data' es una lista de detalles de juegos a guardar, 'filePath' es la ubicación del archivo CSV.
// Retorna un posible error si ocurre durante la escritura del archivo.
func (s *SteamAPI) SaveToCSV(data []steamapi.AppDetails, filePath string) error {
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
			"isFree",
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
				strconv.FormatBool(app.Platforms.Windows),
				strconv.FormatBool(app.Platforms.Mac),
				strconv.FormatBool(app.Platforms.Linux),
				app.ReleaseDate.Date,
				strconv.FormatBool(app.ReleaseDate.ComingSoon),
				app.PriceOverview.Currency,
				strconv.Itoa(int(app.PriceOverview.DiscountPercent)),
				formatInitial(float64(app.PriceOverview.Initial) / 100),
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

// loadExistingData carga los appIDs previamente existentes desde un archivo CSV.
// 'filePath' es la ubicación del archivo CSV.
// Retorna un mapa de appIDs existentes y un posible error si ocurre durante la lectura del archivo.
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

// formatInitial formatea un valor inicial en moneda argentina.
// 'initial' es el valor inicial a formatear.
// Retorna el valor formateado en formato 'ARS X.YY'.
func formatInitial(initial float64) string {
	return fmt.Sprintf("ARS %.2f", initial)
}
