package controller

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func generateUniqueKeyName() string {
	// Obtener un valor de tiempo actual único en nanosegundos
	timestamp := time.Now().UnixNano()

	// Generar un valor aleatorio entre 1 y 1000
	randomValue := rand.Intn(1000)

	// Construir el nombre de la clave único combinando tiempo y valor aleatorio
	uniqueKeyName := fmt.Sprintf("clave%d%d", timestamp, randomValue)

	return uniqueKeyName
}

func createAPIKey() (string, error) {
	// Generar un nombre de clave único
	keyName := generateUniqueKeyName()

	// Comando para crear la clave de API
	cmd := exec.Command("gcloud", "beta", "services", "api-keys", "create", "--display-name="+keyName, "--api-target=service=myapi", "--format=json")

	// Capturar la salida estándar
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Error al ejecutar el comando: %v", err)
	}

	// Convertir la salida a una cadena
	outputStr := string(output)

	// Buscar la clave de API en la salida
	keyIndex := strings.Index(outputStr, "\"keyString\":\"")
	if keyIndex != -1 {
		startIndex := keyIndex + len("\"keyString\":\"")
		endIndex := strings.Index(outputStr[startIndex:], "\"")
		if endIndex != -1 {
			apiKey := outputStr[startIndex : startIndex+endIndex]
			return apiKey, nil
		}
	}

	return "", fmt.Errorf("No se pudo encontrar la clave de API en la salida.")
}
