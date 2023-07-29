package funcgcp

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"steamAPI/api/handlers"
	"steamAPI/api/utilities"
	"time"
)

func ProcessSteamDataAndSaveToStorage(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Método no permitido. Debes utilizar POST para procesar los datos.",
		})
		return
	}

	dataFetcher := &handlers.RealDataFetcher{}

	data, err := dataFetcher.GetData()
	if err != nil {
		log.Printf("Error al obtener los datos de la API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener los datos de la API",
		})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Nombre de archivo no válido",
		})
		return
	}

	// Cargar el archivo CSV completo en Cloud Storage
	err = utilities.UploadFileToGCS(csvContent.String(), "steam-analytics", fileName)
	if err != nil {
		log.Printf("Error al subir el archivo .csv a Cloud Storage: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al subir el archivo .csv a Cloud Storage",
		})
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

	c.JSON(http.StatusOK, response)
}

func isValidFileName(fileName string) bool {
	// Verificar que el nombre cumpla con los requisitos usando expresiones regulares
	// Expresión regular: debe comenzar y terminar con letra o número, y puede contener letras, números, guiones, guiones bajos y puntos
	// Además, debe tener entre 3 y 63 caracteres.
	validNameRegex := regexp.MustCompile(`^[a-z0-9][-a-z0-9_.]{1,61}[a-z0-9]$`)
	return validNameRegex.MatchString(fileName)
}
