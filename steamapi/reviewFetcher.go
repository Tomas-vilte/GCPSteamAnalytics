package steamapi

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

type SteamReviewAPI struct {
	Client HTTPClient
}

// GetReviews GetPositiveReviews GetReviews obtiene las reseñas de un juego específico utilizando su appID.
// Acepta el appID del juego como argumento y devuelve un puntero a la estructura ReviewResponse
// que contiene la información de las reseñas, así como un posible error si ocurre.
func (s *SteamReviewAPI) GetReviews(appID int) (*models.ReviewResponse, error) {
	// Construir la URL de la API de reseñas de Steam
	url := fmt.Sprintf("https://store.steampowered.com/appreviews/%d?json=1&language=latam&filter=recent&num_per_page=100&review_type=positive", appID)

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

func (s *SteamReviewAPI) SaveReviewsToCSV(appID int, reviews *models.ReviewResponse, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Hubo un error al abrir el archivo csv: %v", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		header := []string{
			"SteamAppID",
			"numReviews",
			"AuthorSteamID",
			"NumGamesOwned",
			"NumReviews",
			"PlaytimeForever",
			"PlaytimeLastTwoWeeks",
			"PlaytimeAtReview",
			"LastPlayed",
			"Language",
			"ReviewText",
			"TimestampCreated",
			"VotedUp",
			"VotesUp",
			"VotesFunny",
			"CommentCount",
			"SteamPurchase",
			"ReceivedForFree",
			"WrittenDuringEarlyAccess",
		}
		if err := writer.Write(header); err != nil {
			return err
		}
	}

	for _, review := range reviews.Reviews {
		record := []string{
			strconv.Itoa(appID),
			strconv.Itoa(reviews.ReviewSummary.NumReviews),
			review.Author.SteamID,
			strconv.Itoa(review.Author.NumGamesOwned),
			strconv.Itoa(review.Author.NumReviews),
			strconv.Itoa(review.Author.PlaytimeForever),
			strconv.Itoa(review.Author.PlaytimeLastTwoWeeks),
			strconv.Itoa(review.Author.PlaytimeAtReview),
			strconv.Itoa(review.Author.LastPlayed),
			review.Language,
			review.ReviewText,
			strconv.Itoa(review.TimestampCreated),
			strconv.FormatBool(review.VotedUp),
			strconv.Itoa(review.VotesUp),
			strconv.Itoa(review.VotesFunny),
			strconv.Itoa(review.CommentCount),
			strconv.FormatBool(review.SteamPurchase),
			strconv.FormatBool(review.ReceivedForFree),
			strconv.FormatBool(review.WrittenDuringEarlyAccess),
		}
		if err := writer.Write(record); err != nil {
			log.Printf("Error al escribir en el CSV: %v", err)
			return err
		}
	}

	return nil
}
