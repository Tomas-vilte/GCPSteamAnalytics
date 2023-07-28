package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"steamAPI/api/handlers"
	"steamAPI/api/utilities"
)

func ProcessSteamDataAndSaveToStorage(w http.ResponseWriter, r *http.Request) {

	dataFetcher := &handlers.RealDataFetcher{}

	data, err := dataFetcher.GetData()
	if err != nil {
		log.Printf("Error al obtener los datos de la API: %v", err)
		http.Error(w, "Error al obtener los datos de la API", http.StatusInternalServerError)
		return
	}
	csvContent := "appid,name\n"
	for _, item := range data {
		csvContent += fmt.Sprintf("%d, %s\n", item.Appid, item.Name)
	}

	err = utilities.UploadFileToGCS(csvContent, "steam-analytics", "steam-appids.csv")
	if err != nil {
		log.Printf("Error al subir el archivo .csv a Cloud Storage: %v", err)
		return
	}

	response := map[string]string{
		"message": "Datos obtenidos de la API y guardados en Cloud Storage con exito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
