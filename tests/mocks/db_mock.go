package mocks

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetAllFrom(limit int) ([]entity.Item, error) {
	args := m.Called(limit)
	return args.Get(0).([]entity.Item), args.Error(1)
}

func (m *MockStorage) Update(item entity.Item) error {
	args := m.Called(item)
	return args.Error(0)
}
