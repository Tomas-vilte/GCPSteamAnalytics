package service

import (
	"encoding/json"
	"fmt"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"log"
	"net/http"
)

const (
	baseURL  = "https://store.steampowered.com/api/appdetails"
	apiKey   = "1A059D89640D054BB20FF254FB529E14"
	language = "spanish"
	cc       = "AR"
)

type SteamClient interface {
	GetAppDetails(id int) (map[string]steamapi.AppDetailsResponse, error)
}

func NewSteamClient(client http.Client) SteamClient {
	return &steamClient{
		client: client,
	}
}

type steamClient struct {
	client http.Client
}

func (s *steamClient) GetAppDetails(id int) (map[string]steamapi.AppDetailsResponse, error) {
	url := fmt.Sprintf("%s?l=%s&appids=%d&key=%s&cc=%s", baseURL, language, id, apiKey, cc)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error al crear la solicitud HTTP: %v\n", err)
		return nil, err
	}

	req.Close = true
	response, err := s.client.Do(req)
	if err != nil {
		log.Printf("Error al realizar la solicitud HTTP: %v\n", err)
		return nil, err
	}
	defer response.Body.Close()
	var responseData map[string]steamapi.AppDetailsResponse
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		log.Printf("Error al decodificar la respuesta JSON: %v\n", err)
		return nil, err
	}

	return responseData, nil
}
