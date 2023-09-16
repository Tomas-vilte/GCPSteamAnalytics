package tests

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/Tomas-vilte/GCPSteamAnalytics/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
            "description": "A popular first-person shooter game.",
			"release_date": {
                "coming_soon": false,
                "date": "21 AGO 2012"
            }
        }
    }
}`
	responseBytes := []byte(responseJSON)
	mockGames := []entity.Item{
		{Appid: 730,
			Name:      "Counter-Strike: Global Offensive",
			Status:    "PENDING",
			IsValid:   false,
			CreatedAt: &time.Time{},
			UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true}},
	}

	appDetails, err := processor.ProcessResponse([][]byte{responseBytes}, mockGames)
	assert.NoError(t, err)
	assert.NotNil(t, appDetails)

	mockStorage.AssertExpectations(t)
	mockSteamClient.AssertExpectations(t)
}

func TestGameProcessor_UpdateData(t *testing.T) {
	games := []entity.Item{
		{Appid: 730, IsValid: false},
	}
	id := int64(730)
	isValid := true

	storageMock := &mocks.MockStorage{}
	storageMock.On("Update", mock.Anything).Return(nil)

	processor := service.NewGameProcessor(storageMock, nil)

	err := processor.UpdateData(games, id, isValid)

	assert.NoError(t, err, "Se esperaba que no haya error")
	updatedGame := games[0]
	assert.True(t, updatedGame.IsValid, "El estado del juego no se actualizó correctamente")
	storageMock.AssertCalled(t, "Update", updatedGame)
}
