package GCPSteamAnalytics

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Tomas-vilte/GCPSteamAnalytics/functionGCP"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func OkResponse(w http.ResponseWriter, r *http.Request) {
	// Respuesta que indica que todo está ok.
	message := "Todo ok"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: message})
}

func init() {
	// Registra la función HTTP con el nombre "ProcessSteamDataAndSaveToStorage"
	functions.HTTP("ProcessSteamDataAndSaveToStorage", functionGCP.ProcessSteamDataAndSaveToStorage)
	// Registra la función HTTP con el nombre "OkResponse" para el endpoint /test
	functions.HTTP("OkResponse", OkResponse)
}
