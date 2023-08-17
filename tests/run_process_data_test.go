package tests

import (
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockSteamAPI struct {
	mock.Mock
}

func (m *MockSteamAPI) GetStartIndexToProcess(lastProcessedAppID int, appIDs []int) int {
	//TODO implement me
	panic("implement me")
}

func (m *MockSteamAPI) IsEmptyAppID(appID int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSteamAPI) AddToEmptyAppIDsTable(appID int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSteamAPI) ProcessAppID(id int) (*models.AppDetails, error) {
	args := m.Called(id)
	return args.Get(0).(*models.AppDetails), args.Error(1)
}

func (m *MockSteamAPI) ProcessSteamData(appIDs []int, limit int) ([]models.AppDetails, error) {
	args := m.Called(appIDs, limit)
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
	mockAPI := new(MockSteamAPI)
	lastProcessedAppID := 0
	appIDs := []int{1, 2, 3}
	data := []models.AppDetails{{}, {}, {}}

	mockAPI.On("LoadLastProcessedAppid").Return(lastProcessedAppID, nil)
	mockAPI.On("GetAllAppIDs", lastProcessedAppID).Return(appIDs, nil)
	mockAPI.On("ProcessSteamData", appIDs, 10).Return(data, nil)
	mockAPI.On("SaveToCSV", data, "../data/dataDetails.csv").Return(nil)
	err := steamapi.RunProcessData(mockAPI, 10)

	assert.NoError(t, err)
	mockAPI.AssertExpectations(t)
}
