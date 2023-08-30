package tests

import (
	"context"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/Tomas-vilte/GCPSteamAnalytics/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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
