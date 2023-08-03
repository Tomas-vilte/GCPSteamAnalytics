package GCPSteamAnalytics

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Tomas-vilte/GCPSteamAnalytics/functionGCP"
	"net/http"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/test-function":
		functionGCP.CheckHealth(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/getGamesFromSteam":
		functionGCP.ProcessSteamDataAndSaveToStorage(w, r)
	default:
		http.Error(w, "Ruta inválida.", http.StatusNotFound)
	}
}

func init() {
	functions.HTTP("GameHandler", GameHandler)
}
