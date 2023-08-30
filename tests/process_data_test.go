package tests

import (
	"context"
	"database/sql"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/Tomas-vilte/GCPSteamAnalytics/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGameProcessor_GetGamesFromAPI(t *testing.T) {
	mockSteamClient := new(mocks.MockSteamClient)
	mockStorage := new(mocks.MockStorage)

	mockResponse := []byte(`{"data": {"name": "Test Game"}}`)
	mockSteamClient.On("GetAppDetails", mock.Anything).Return(mockResponse, nil)

	processor := service.NewGameProcessor(mockStorage, mockSteamClient)

	mockItems := []entity.Item{
		{Appid: 123},
		{Appid: 456},
	}

	ctx := context.Background()
	responseData, err := processor.GetGamesFromAPI(ctx, mockItems)
	assert.NoError(t, err)
	assert.NotNil(t, responseData)
	assert.Len(t, responseData, len(mockItems))

	mockSteamClient.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestGameProcessor_GetGamesFromAPI_ErrorInGetAppDetails(t *testing.T) {
	mockSteamClient := new(mocks.MockSteamClient)
	mockStorage := new(mocks.MockStorage)

	mockResponse := []byte(`{"data": {"name": "Test Game"}}`)
	mockSteamClient.On("GetAppDetails", mock.Anything).Return(mockResponse, nil)

	processor := service.NewGameProcessor(mockStorage, mockSteamClient)

	mockItems := []entity.Item{
		{Appid: 123},
		{Appid: 456},
	}

	ctx := context.Background()
	responseData, err := processor.GetGamesFromAPI(ctx, mockItems)
	assert.NoError(t, err)
	assert.NotNil(t, responseData)
	assert.Len(t, responseData, len(mockItems))

	mockSteamClient.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestGameProcessor_ProcessResponse(t *testing.T) {
	mockStorage := new(mocks.MockStorage)
	mockSteamClient := new(mocks.MockSteamClient)

	processor := service.NewGameProcessor(mockStorage, mockSteamClient)
	mockStorage.On("Update", mock.Anything).Return(nil)

	responseJSON := `
   {
    "730": {
        "success": true,
        "data": {
            "type": "game",
            "name": "Counter-Strike: Global Offensive",
            "description": "A popular first-person shooter game."
        }
    }
}`
	responseBytes := []byte(responseJSON)
	mockGames := []entity.Item{
		{Appid: 730,
			Name:      "Counter-Strike: Global Offensive",
			Status:    "PENDING",
			IsValid:   false,
			CreatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true}},
	}

	appDetails, err := processor.ProcessResponse([][]byte{responseBytes}, mockGames)
	assert.NoError(t, err)
	assert.NotNil(t, appDetails)

	mockStorage.AssertExpectations(t)
	mockSteamClient.AssertExpectations(t)
}
