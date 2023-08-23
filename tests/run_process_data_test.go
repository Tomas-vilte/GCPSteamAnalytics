package tests

import (
	"context"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockSteamAPI struct {
	mock.Mock
}

func (m *MockSteamAPI) GetStartIndexToProcess(lastProcessedAppID int, appIDs []int) int {
	args := m.Called(lastProcessedAppID, appIDs)
	return args.Int(0)
}

func (m *MockSteamAPI) AreEmptyAppIDs(appIDs []int) (map[int]bool, error) {
	resultMap := make(map[int]bool)
	for _, id := range appIDs {
		args := m.Called(id)
		resultMap[id] = args.Bool(0) // Suponiendo que el resultado está en la posición 0
	}
	return resultMap, nil // Simulamos que no hay error
}

func (m *MockSteamAPI) AddToEmptyAppIDsTable(appID int) error {
	args := m.Called(appID)
	return args.Error(0)
}

func (m *MockSteamAPI) ProcessAppID(id int) (*models.AppDetails, error) {
	args := m.Called(id)
	return args.Get(0).(*models.AppDetails), args.Error(1)
}

func (m *MockSteamAPI) ProcessSteamData(ctx context.Context, appIDs []int, limit int) ([]models.AppDetails, error) {
	args := m.Called(ctx, appIDs, limit)
	return args.Get(0).([]models.AppDetails), args.Error(1)
}

func (m *MockSteamAPI) GetAllAppIDs(lastProcessedAppID int) ([]int, error) {
	args := m.Called(lastProcessedAppID)
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockSteamAPI) LoadLastProcessedAppid() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockSteamAPI) SaveLastProcessedAppid(lastProcessedAppid int) error {
	args := m.Called(lastProcessedAppid)
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
	lastProcessedAppID := 100
	appIDs := []int{101, 102, 103}
	startIndex := 0
	data := []models.AppDetails{{SteamAppid: 101}, {SteamAppid: 102}, {SteamAppid: 103}}
	limit := 10

	mockSteamData.On("LoadLastProcessedAppid").Return(lastProcessedAppID, nil)
	mockSteamData.On("GetAllAppIDs", lastProcessedAppID).Return(appIDs, nil)
	mockSteamData.On("GetStartIndexToProcess", lastProcessedAppID, appIDs).Return(startIndex)
	mockSteamData.On("ProcessSteamData", ctx, appIDs[startIndex:], limit).Return(data, nil)
	mockSteamData.On("SaveToCSV", data, mock.AnythingOfType("string")).Return(nil)
	// Llama a la función que deseas probar
	err := steamapi.RunProcessData(mockSteamData, limit)

	// Verifica que no haya error y que los métodos del mock hayan sido llamados
	assert.NoError(t, err)
	mockSteamData.AssertExpectations(t)
}
