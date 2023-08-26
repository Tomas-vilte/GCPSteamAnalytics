package tests

import (
	"context"
	"testing"

	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSteamAPI struct {
	mock.Mock
}

func (m *MockSteamAPI) ProcessAppID(id int) (*models.AppDetails, error) {
	args := m.Called(id)
	return args.Get(0).(*models.AppDetails), args.Error(1)
}

func (m *MockSteamAPI) ProcessSteamData(ctx context.Context, appIDs []int, limit int) ([]models.AppDetails, error) {
	args := m.Called(ctx, appIDs, limit)
	return args.Get(0).([]models.AppDetails), args.Error(1)
}

func (m *MockSteamAPI) GetAllAppIDs(limit int) ([]int, error) {
	args := m.Called(limit)
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockSteamAPI) UpdateAppStatus(id int, isValid bool) error {
	args := m.Called(id, isValid)
	return args.Error(0)
}

func (m *MockSteamAPI) SaveToCSV(data []models.AppDetails, filePath string) error {
	args := m.Called(data, filePath)
	return args.Error(0)
}

func TestRunProcessData(t *testing.T) {
	// Crea una instancia del mock
	mockSteamData := new(MockSteamAPI)
	ctx := context.Background()
	// Configura el comportamiento del mock
	limit := 100
	appIDs := []int{101, 102, 103}
	data := []models.AppDetails{{SteamAppid: 101}, {SteamAppid: 102}, {SteamAppid: 103}}

	mockSteamData.On("GetAllAppIDs", limit).Return(appIDs, nil)
	mockSteamData.On("ProcessSteamData", ctx, appIDs, limit).Return(data, nil)
	mockSteamData.On("SaveToCSV", data, mock.AnythingOfType("string")).Return(nil)
	// Llama a la función que deseas probar
	err := steamapi.RunProcessData(mockSteamData, limit)

	// Verifica que no haya error y que los métodos del mock hayan sido llamados
	assert.NoError(t, err)
	mockSteamData.AssertExpectations(t)
}
