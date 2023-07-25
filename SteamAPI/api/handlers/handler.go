package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// DataFetcher representa la interfaz para obtener datos.
type DataFetcher interface {
	GetData() ([]Item, error)
}

// RealDataFetcher implementa DataFetcher para obtener datos reales de la API.
type RealDataFetcher struct{}

// APIResponse representa la estructura del JSON devuelto por la API.
type APIResponse struct {
	Data struct {
		StoreItems []Item `json:"store_items"`
	} `json:"data"`
}

// Item representa cada elemento del array "store_items".
type Item struct {
	Appid int64  `json:"appid"`
	Name  string `json:"name"`
}

// GetData realiza una solicitud HTTP GET al endpoint "http://localhost:5000/datos" para obtener los datos.
// Retorna una lista de elementos (Item) y un error en caso de que la solicitud falle o el JSON no pueda ser decodificado.
func (r *RealDataFetcher) GetData() ([]Item, error) {
	// Realizar la solicitud HTTP GET a la API para obtener los datos.
	response, err := http.Get("http://localhost:5000/datos")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Decodificar la respuesta JSON en la estructura APIResponse.
	var apiResponse APIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse.Data.StoreItems, nil
}

// Definir un error personalizado para cuando falle la obtención de datos.
var ErrDataFetch = errors.New("error al obtener datos")
