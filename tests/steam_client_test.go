package tests

import (
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/Tomas-vilte/GCPSteamAnalytics/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestSteamClient_GetAppDetails_Success(t *testing.T) {
	mockResponse := `
	{
		"12345": {
			"data": {
				"name": "Test Game",
				"price_overview": {
					"initial": 999,
					"final": 499
				}
			}
		}
	}`
	mockHTTPClient := new(mocks.MockSteamClient)
	mockHTTPClient.On("Do", mock.Anything).Return(
		&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(mockResponse)),
		},
		nil,
	)

	client := service.NewSteamClient(mockHTTPClient)

	responseData, err := client.GetAppDetails(12345)
	assert.NoError(t, err)
	assert.NotNil(t, responseData)

	expectedResponse := []byte(mockResponse)
	assert.Equal(t, expectedResponse, responseData)

	mockHTTPClient.AssertExpectations(t)
}

func TestSteamClient_GetAppDetails_ErrorRequest(t *testing.T) {
	mockHTTPClient := new(mocks.MockSteamClient)
	mockHTTPClient.On("Do", mock.Anything).Return(
		nil,
		errors.New("error making request"),
	)

	client := service.NewSteamClient(mockHTTPClient)

	responseData, err := client.GetAppDetails(12345)
	assert.Error(t, err)
	assert.Nil(t, responseData)

	mockHTTPClient.AssertExpectations(t)
}
