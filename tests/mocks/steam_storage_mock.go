package mocks

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetGamesByPage(filter string, startIndex, pageSize int) ([]entity.GameDetails, int, error) {
	argsList := m.Called(filter, startIndex, pageSize)
	return argsList.Get(0).([]entity.GameDetails), argsList.Int(1), argsList.Error(2)
}

func (m *MockStorage) GetAllFrom(limit int) ([]entity.Item, error) {
	args := m.Called(limit)
	return args.Get(0).([]entity.Item), args.Error(1)
}

func (m *MockStorage) Update(item entity.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockStorage) GetAllByAppID(appID int) ([]entity.Item, error) {
	args := m.Called(appID)
	return args.Get(0).([]entity.Item), args.Error(1)

}

func (m *MockStorage) GetGameDetails(gameID int) (*entity.GameDetails, error) {
	args := m.Called(gameID)
	return args.Get(0).(*entity.GameDetails), args.Error(1)
}

func (m *MockStorage) SaveGameDetails(dataProcessed []model.AppDetails) error {
	args := m.Called(dataProcessed)
	return args.Error(0)
}
