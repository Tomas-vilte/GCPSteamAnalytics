package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetSupportedLanguagesString convierte un slice de idiomas soportados en un string formateado.
// Si no hay idiomas soportados, retorna un mensaje indicando que no hay soporte.
func GetSupportedLanguagesString(supportedLanguages []string) string {
	if len(supportedLanguages) == 0 {
		return "No hay soporte para este tipo de idioma"
	}
	return strings.Join(supportedLanguages, ", ")
}

// ParseSupportedLanguages analiza una cadena raw de idiomas soportados y la convierte en un mapa.
func ParseSupportedLanguages(raw string) map[string][]string {
	// Creamos un mapa para almacenar los idiomas soportados por cada tipo.
	languages := make(map[string][]string)

	// Divide la cadena raw en partes separadas por ", ".
	parts := strings.Split(raw, ", ")
	for _, part := range parts {
		if strings.HasSuffix(part, "<strong>*</strong>") {
			// Si la parte termina con "<strong>*</strong>", se trata de un idioma con soporte completo.
			// Eliminar "<strong>*</strong>" del final para obtener el nombre del idioma.
			lang := strings.TrimSuffix(part, "<strong>*</strong>")

			// Agrega el idioma a las listas de idiomas para cada tipo.
			languages["full_audio"] = append(languages["full_audio"], lang)
			languages["interface"] = append(languages["interface"], lang)
			languages["subtitles"] = append(languages["subtitles"], lang)
		} else {
			// Si la parte no termina con "<strong>*</strong>", se trata de un idioma sin soporte de audio.
			// Elimina "<br><strong>*</strong>idiomas con localización de audio" del final para obtener el nombre del idioma.
			lang := strings.TrimSuffix(part, "<br><strong>*</strong>idiomas con localización de audio")

			// Agrega el idioma a las listas de idiomas para los tipos "interface" y "subtitles".
			languages["interface"] = append(languages["interface"], lang)
			languages["subtitles"] = append(languages["subtitles"], lang)
		}
	}

	return languages
}

// FormatInitial formatea un valor inicial en moneda argentina.
// 'initial' es el valor inicial a formatear.
// Retorna el valor formateado en formato 'ARS X.YY'.
func FormatInitial(initial float64) string {
	return fmt.Sprintf("ARS %.2f", initial)
}

// LoadExistingData carga los appIDs previamente existentes desde un archivo CSV.
// 'filePath' es la ubicación del archivo CSV.
// Retorna un mapa de appIDs existentes y un posible error si ocurre durante la lectura del archivo.
func LoadExistingData(filePath string) (map[int]bool, error) {
	existingData := make(map[int]bool)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Si el archivo no existe, devolver el mapa vacío sin error.
			log.Printf("El archivo no existe: %v\n", err)
			return existingData, nil
		}
		log.Printf("Error al abrir el archivo: %v\n", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Leer y descartar la primera fila (encabezados)
	_, err = reader.Read()
	if err != nil {
		log.Printf("Error al leer la primera fila: %v\n", err)
		return nil, err
	}

	// Leer las filas restantes y procesar los appIDs
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error al leer fila: %v\n", err)
			return nil, err
		}

		// Asegurarse de que haya al menos un valor en el registro antes de convertir
		if len(record) < 1 {
			log.Printf("Registro sin valores, saltando...\n")
			continue
		}

		appID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error al convertir appID a entero: %v\n", err)
			continue // Saltar esta fila y seguir con la siguiente
		}
		existingData[appID] = true
	}

	log.Printf("Carga de datos existentes completada. Total de appIDs cargados: %d\n", len(existingData))
	return existingData, nil
}

func ReadAppIDsFromCSV(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error al abrir el archivo: %v", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read()

	var appIDs []int

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error al leer la primera fila: %v\n", err)
			return nil, err
		}
		appID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error al convertir appID a entero: %v\n", err)
			return nil, err
		}
		appIDs = append(appIDs, appID)
	}
	return appIDs, nil
}
