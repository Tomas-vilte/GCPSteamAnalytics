package GCPSteamAnalytics

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Tomas-vilte/GCPSteamAnalytics/functionGCP"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	// Definir el manejador para la ruta /test con el método GET
	r.HandleFunc("/test", OkResponse).Methods("GET")

	// Definir el manejador para la ruta /dbgames con el método POST
	r.HandleFunc("/dbgames", functionGCP.ProcessSteamDataAndSaveToStorage).Methods("POST")

	// Registrar el enrutador en la función HTTP
	functions.HTTP("api", r.ServeHTTP)
}
