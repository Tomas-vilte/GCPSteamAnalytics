package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type CreateApiKey interface {
	CreateKey(ctx *gin.Context)
}

func CreateKey(ctx *gin.Context) {
	apiKey, err := createAPIKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Clave api creada con exito.",
		"api_key": apiKey,
	})
}

func generateUniqueKeyName() string {
	timestamp := time.Now().UnixNano()

	randomValue := rand.Intn(1000)

	uniqueKeyName := fmt.Sprintf("clave%d%d", timestamp, randomValue)

	return uniqueKeyName
}

func createAPIKey() (string, error) {
	keyName := generateUniqueKeyName()

	cmd := exec.Command("gcloud", "beta", "services", "api-keys", "create", "--display-name="+keyName, "--api-target=service=myapi", "--format=json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error al ejecutar el comando: %v", err)
	}

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

	return "", fmt.Errorf("no se pudo encontrar la clave de API en la salida")
}
