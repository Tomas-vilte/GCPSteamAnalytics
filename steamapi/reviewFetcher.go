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

func (s *SteamReviewAPI) GetReviews(appID int) (*models.ReviewResponse, error) {
	url := fmt.Sprintf("https://store.steampowered.com/appreviews/%d?json=1&language=latam&filter=recent&num_per_page=50", appID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error al hacer una peticion a la API: %v", err)
		return nil, err
	}
	req.Close = true
	response, err := s.Client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	defer response.Body.Close()

	var reviewResponse models.ReviewResponse
	err = json.NewDecoder(response.Body).Decode(&reviewResponse)
	if err != nil {
		log.Printf("Error a decodificar la respuesta: %v\n", err)
		return nil, err
	}
	return &reviewResponse, nil
}
