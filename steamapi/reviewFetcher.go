package steamapi

import (
	"encoding/json"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"log"
	"net/http"
)

type SteamReviewAPI struct {
	Client HTTPClient
}

// GetReviews obtiene las reseñas de un juego específico utilizando su appID.
// Acepta el appID del juego como argumento y devuelve un puntero a la estructura ReviewResponse
// que contiene la información de las reseñas, así como un posible error si ocurre.
func (s *SteamReviewAPI) GetReviews(appID int) (*models.ReviewResponse, error) {
	// Construir la URL de la API de reseñas de Steam
	url := fmt.Sprintf("https://store.steampowered.com/appreviews/%d?json=1&language=latam&filter=recent&num_per_page=50", appID)

	// Crear una nueva solicitud HTTP GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error al crear la solicitud HTTP: %v\n", err)
		return nil, err
	}
	req.Close = true

	// Enviar la solicitud y obtener la respuesta
	response, err := s.Client.Do(req)
	if err != nil {
		log.Printf("Error al realizar la solicitud HTTP: %v\n", err)
		return nil, err
	}
	defer response.Body.Close()

	// Decodificar la respuesta JSON en la estructura ReviewResponse
	var reviewResponse models.ReviewResponse
	err = json.NewDecoder(response.Body).Decode(&reviewResponse)
	if err != nil {
		log.Printf("Error a decodificar la respuesta: %v\n", err)
		return nil, err
	}

	// Devolver un puntero a la estructura ReviewResponse y un posible error
	return &reviewResponse, nil
}
