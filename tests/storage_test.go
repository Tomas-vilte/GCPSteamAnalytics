package tests

import (
	"database/sql"
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/tests/mocks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestStorage_GetAllFrom_Success(t *testing.T) {
	mockItems := []entity.Item{
		{
			Appid:     123,
			Name:      "Test Game",
			Status:    "PENDING",
			IsValid:   false,
			CreatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
		},
	}

	mockStorage := new(mocks.MockStorage)
	mockStorage.On("GetAllFrom", mock.Anything).Return(mockItems, nil)

	storage := mockStorage

	limit := 10
	items, err := storage.GetAllFrom(limit)
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Len(t, items, len(mockItems))

	mockStorage.AssertExpectations(t)
}

func TestStorage_GetAllFrom_Error(t *testing.T) {
	mockError := errors.New("mock error")
	var mockItems []entity.Item

	mockStorage := new(mocks.MockStorage)
	mockStorage.On("GetAllFrom", mock.Anything).Return(mockItems, mockError)

	storage := mockStorage

	limit := 10
	items, err := storage.GetAllFrom(limit)
	assert.Error(t, err)
	assert.Nil(t, items)
	assert.Equal(t, mockError, err)

	mockStorage.AssertExpectations(t)
}

func TestStorage_Update_Success(t *testing.T) {
	mockStorage := new(mocks.MockStorage)
	mockStorage.On("Update", mock.Anything).Return(nil)

	storage := mockStorage

	mockItem := entity.Item{
		Appid:     123,
		Name:      "Test Game",
		Status:    "PENDING",
		IsValid:   false,
		CreatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := storage.Update(mockItem)
	assert.NoError(t, err)

	mockStorage.AssertExpectations(t)
}

func TestStorage_Update_Error(t *testing.T) {
	mockError := errors.New("mock error")

	mockStorage := new(mocks.MockStorage)
	mockStorage.On("Update", mock.Anything).Return(mockError)

	storage := mockStorage

	mockItem := entity.Item{
		Appid:     123,
		Name:      "Test Game",
		Status:    "PENDING",
		IsValid:   false,
		CreatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := storage.Update(mockItem)
	assert.Error(t, err)
	assert.Equal(t, mockError, err)

	mockStorage.AssertExpectations(t)
}

func TestStorage_Update_ItemNotFound(t *testing.T) {
	mockError := errors.New("item not found")

	mockStorage := new(mocks.MockStorage)
	mockStorage.On("Update", mock.Anything).Return(mockError)

	storage := mockStorage

	mockItem := entity.Item{
		Appid:     123,
		Name:      "Test Game",
		Status:    "PENDING",
		IsValid:   false,
		CreatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := storage.Update(mockItem)
	assert.Error(t, err)
	assert.Equal(t, mockError, err)

	mockStorage.AssertExpectations(t)
}
