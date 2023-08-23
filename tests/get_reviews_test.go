package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
	mock.Mock
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetReviews(t *testing.T) {
	MockResponse := models.ReviewResponse{
		Success: 1,
		ReviewSummary: models.ReviewSummary{
			NumReviews:      50,
			ReviewScore:     8,
			ReviewScoreDesc: "Very Positive",
			TotalPositive:   6488,
			TotalNegative:   837,
			TotalReviews:    7325,
		},
		Reviews: []models.Review{
			{
				RecommendationID: "144456675",
				Author: models.ReviewAuthor{
					SteamID:              "76561199263188792",
					NumGamesOwned:        21,
					NumReviews:           1,
					PlaytimeForever:      72380,
					PlaytimeLastTwoWeeks: 3377,
					PlaytimeAtReview:     72303,
					LastPlayed:           1692330585,
				},
				Language:                 "latam",
				ReviewText:               "juegaso\n",
				TimestampCreated:         1692326452,
				TimestampUpdated:         1692326452,
				VotedUp:                  true,
				VotesUp:                  0,
				VotesFunny:               0,
				CommentCount:             0,
				SteamPurchase:            true,
				ReceivedForFree:          false,
				WrittenDuringEarlyAccess: false,
			},
		},
	}

	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Simula una respuesta exitosa
			mockResponseBytes, _ := json.Marshal(MockResponse)
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(mockResponseBytes)),
			}, nil
		},
	}

	api := &steamapi.SteamReviewAPI{
		Client: mockClient,
	}

	appID := 730
	reviewResponse, err := api.GetReviews(appID)
	if err != nil {
		t.Errorf("Error inesperado: %v", err)
	}
	// Verifica que la respuesta coincida con el mockResponse
	if !reflect.DeepEqual(reviewResponse, &MockResponse) {
		t.Errorf("La respuesta no coincide con el mockResponse")
	}

	assert.NoError(t, err)
	assert.NotNil(t, reviewResponse)
	assert.Equal(t, MockResponse, *reviewResponse)
}

func TestGetReviews_HTTPError(t *testing.T) {
	// Creamos un cliente de prueba que devuelve un error HTTP
	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, fmt.Errorf("error de red")
		},
	}

	// Creamos una instancia de SteamReviewAPI con el cliente de prueba
	api := &steamapi.SteamReviewAPI{
		Client: mockClient,
	}

	// Llama a la función GetReviews con un appID ficticio
	_, err := api.GetReviews(730)
	if err == nil {
		t.Errorf("Se esperaba un error, pero no se recibió ninguno")
		return
	}

	/// Comprueba si el error coincide con el error esperado
	expectedError := "error de red"
	if strings.ToLower(err.Error()) != strings.ToLower(expectedError) {
		t.Errorf("Error incorrecto. Se esperaba %s, pero se obtuvo %s", expectedError, err.Error())
	}
}

func TestGetReviews_JSONDecodeError(t *testing.T) {
	invalidJSON := `{
   "success":1,
   "query_summary":{
      "num_reviews":50,
      "review_score":8,
      "review_score_desc":"Very Positive",
      "total_positive":6488,
      "total_negative":837,
      "total_reviews":7325
   },
   "reviews":[
      {
         "recommendationid":"144456675"
      }
   ],
}`

	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			response := http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(invalidJSON)),
			}
			return &response, nil
		},
	}

	reviewAPI := &steamapi.SteamReviewAPI{
		Client: mockClient,
	}

	_, err := reviewAPI.GetReviews(730)
	if err == nil {
		t.Errorf("Se esperaba un error al decodificar la respuesta JSON inválida")
	}
}
