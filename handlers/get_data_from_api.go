package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// DataFetcher representa la interfaz para obtener datos.
type DataFetcher interface {
	GetData() ([]Item, error)
}

// RealDataFetcher implementa DataFetcher para obtener datos reales de la API.
type RealDataFetcher struct{}

// APIResponse representa la estructura del JSON devuelto por la API.
type APIResponse struct {
	Applist struct {
		Apps []Item `json:"apps"`
	} `json:"applist"`
}

// Item representa cada elemento del array "store_items".
type Item struct {
	Appid     int64  `json:"appid"`
	Name      string `json:"name"`
	Status    string
	IsValid   bool
	CreatedAt time.Time
	UpdateAt  time.Time
}

// GetData realiza una solicitud HTTP GET al endpoint de Steam para obtener los datos.
// Retorna una lista de elementos (Item) y un error en caso de que la solicitud falle o el JSON no pueda ser decodificado.
func (r *RealDataFetcher) GetData() ([]Item, error) {
	// Realizar la solicitud HTTP GET a la API para obtener los datos.
	response, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v0002/?key=1A059D89640D054BB20FF254FB529E14&format=json")
	if err != nil {
		log.Printf("Error al realizar una request: %v", err)
		return nil, err
	}

	defer response.Body.Close()

	// Decodificar la respuesta JSON en la estructura APIResponse.
	var apiResponse APIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		log.Printf("Error al decodificar la respuesta JSON: %v", err)
		return nil, err
	}

	// Quitamos valores vacios de la columna name
	var filteredApps []Item
	for _, app := range apiResponse.Applist.Apps {
		if app.Name != "" {
			item := Item{
				Appid:     app.Appid,
				Name:      app.Name,
				Status:    "Pending",
				IsValid:   false,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			}
			filteredApps = append(filteredApps, item)
		}
	}

	return filteredApps, nil
}

// Definir un error personalizado para cuando falle la obtención de datos.
var ErrDataFetch = errors.New("error al obtener datos")
