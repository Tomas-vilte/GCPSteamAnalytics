package functionGCP

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
)

func ProcessSteamDataAndSaveToStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido. Debes utilizar POST para procesar los datos.", http.StatusMethodNotAllowed)
		return
	}
	dataFetcher := &handlers.RealDataFetcher{}

	data, err := dataFetcher.GetData()
	if err != nil {
		log.Printf("Error al obtener los datos de la API: %v", err)
		http.Error(w, "Error al obtener los datos de la API", http.StatusInternalServerError)
		return
	}

	var csvContent bytes.Buffer
	writer := csv.NewWriter(&csvContent)

	// Escribir encabezados
	headers := []string{"appid", "name"}
	_ = writer.Write(headers)

	// Escribir datos
	for _, item := range data {
		if item.Name != "" {
			row := []string{fmt.Sprintf("%d", item.Appid), item.Name}
			_ = writer.Write(row)
		}
	}

	writer.Flush()

	// Generar el nombre del archivo CSV con la fecha y hora actual
	fileName := fmt.Sprintf("steam_data_%s.csv", time.Now().Format("2006-01-02-15-04-05"))

	// Verificar si el nombre del archivo cumple con los requisitos
	if !isValidFileName(fileName) {
		log.Printf("Nombre de archivo no válido: %s", fileName)
		http.Error(w, "Nombre de archivo no válido", http.StatusInternalServerError)
		return
	}

	// Cargar el archivo CSV completo en Cloud Storage
	err = utils.UploadFileToGCS(csvContent.String(), "steam-analytics", fileName)
	if err != nil {
		log.Printf("Error al subir el archivo .csv a Cloud Storage: %v", err)
		http.Error(w, "Error al subir el archivo .csv a Cloud Storage", http.StatusInternalServerError)
		return
	}

	startTime := time.Now()
	// Calcular la duración del proceso de carga
	duration := time.Since(startTime).Seconds()

	// Fecha y hora en que se subió el archivo
	uploadTime := time.Now().Format(time.RFC3339)

	// Datos adicionales en la respuesta
	response := map[string]interface{}{
		"success":     true,
		"message":     "Datos obtenidos de la API y guardados en Cloud Storage con éxito",
		"timestamp":   time.Now(),
		"data_count":  len(data),
		"file_url":    fmt.Sprintf("https://storage.googleapis.com/steam-analytics/%s", fileName),
		"duration":    duration,
		"upload_time": uploadTime,
	}

	writeJSONResponse(w, response, http.StatusOK)
}

// Verifica que el nombre cumpla con los requisitos usando expresiones regulares
// Expresión regular: debe comenzar y terminar con letra o número, y puede contener letras, números, guiones, guiones bajos y puntos
// Además, debe tener entre 3 y 63 caracteres.
func isValidFileName(fileName string) bool {
	validNameRegex := regexp.MustCompile(`^[a-z0-9][-a-z0-9_.]{1,61}[a-z0-9]$`)
	return validNameRegex.MatchString(fileName)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {

	// Prepara el mensaje de respuesta
	message := fmt.Sprintf("La función ProcessSteamDataAndSaveToStorage está funcionando bien, y su código de estado es %d", http.StatusOK)

	// Código de estado de la función ProcessSteamDataAndSaveToStorage
	statusCode := http.StatusOK

	// Crear el JSON de respuesta
	response := map[string]interface{}{
		"message":     message,
		"status_code": statusCode,
	}

	// Convertir el JSON a bytes
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error al formatear la respuesta JSON", http.StatusInternalServerError)
		return
	}

	// Establecer las cabeceras para la respuesta JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Escribir la respuesta JSON en el http.ResponseWriter
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Error al escribir la respuesta", http.StatusInternalServerError)
		return
	}
}
