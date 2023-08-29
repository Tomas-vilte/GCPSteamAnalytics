package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL  = "https://store.steampowered.com/api/appdetails"
	apiKey   = "1A059D89640D054BB20FF254FB529E14"
	language = "spanish"
	cc       = "AR"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SteamClient interface {
	GetAppDetails(id int) ([]byte, error)
}

type steamClient struct {
	client HTTPClient
}

func NewSteamClient(client HTTPClient) *steamClient {
	return &steamClient{
		client: client,
	}
}

func (s *steamClient) GetAppDetails(id int) ([]byte, error) {
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

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error al leer la respuesta HTTP: %v\n", err)
		return nil, err
	}
	return responseData, nil
}
