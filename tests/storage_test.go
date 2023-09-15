package tests

import (
	"database/sql"
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"

	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/tests/mocks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStorage_GetAllFrom_Success(t *testing.T) {
	mockItems := []entity.Item{
		{
			Appid:     123,
			Name:      "Test Game",
			Status:    "PENDING",
			IsValid:   false,
			CreatedAt: &time.Time{},
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
		CreatedAt: &time.Time{},
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
		CreatedAt: &time.Time{},
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
		CreatedAt: &time.Time{},
		UpdatedAt: &sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := storage.Update(mockItem)
	assert.Error(t, err)
	assert.Equal(t, mockError, err)

	mockStorage.AssertExpectations(t)
}

func TestGetGamesByPage(t *testing.T) {
	mockDB := new(mocks.MockStorage)

	expectedGames := []entity.GameDetails{{ID: 1, Name: "Game 1"}}
	mockDB.On("GetGamesByPage", "game", 0, 10).Return(expectedGames, 1, nil)

	games, totalItems, err := mockDB.GetGamesByPage("game", 0, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedGames, games)
	assert.Equal(t, 1, totalItems)
	mockDB.AssertExpectations(t)
}

func TestSaveGameDetails(t *testing.T) {
	mockDB := new(mocks.MockStorage)

	mockDB.On("SaveGameDetails", mock.AnythingOfType("[]model.AppDetails")).Return(nil)

	appDetails := []model.AppDetails{
		{
			SteamAppid:  1,
			Name:        "Juego 1",
			Description: "Descripción del juego 1",
		},
		{
			SteamAppid:  2,
			Name:        "Juego 2",
			Description: "Descripción del juego 2",
		},
	}

	err := mockDB.SaveGameDetails(appDetails)

	assert.NoError(t, err)

	mockDB.AssertCalled(t, "SaveGameDetails", appDetails)
}

func TestGetGameDetails_Success(t *testing.T) {
	mockDB := new(mocks.MockStorage)
	expectedGameDetails := entity.GameDetails{}
	mockDB.On("QueryRowx", mock.Anything, mock.Anything).Return(mockDB)
	mockDB.On("StructScan", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*entity.GameDetails)
		*dest = expectedGameDetails
	})

	gameID := 730 // ID de juego válido
	mockDB.On("GetGameDetails", gameID).Return(&expectedGameDetails, nil)

	result, err := mockDB.GetGameDetails(gameID)

	require.NoError(t, err)

	assert.Equal(t, &expectedGameDetails, result)
}

func TestGetAllByAppID(t *testing.T) {
	mockStorage := new(mocks.MockStorage)

	expectedItems := []entity.Item{
		{Appid: 1, Name: "Juego1", Status: "Processed", IsValid: true},
		{Appid: 2, Name: "Juego2", Status: "Pending", IsValid: false},
	}

	mockStorage.On("GetAllByAppID", 1).Return(expectedItems, nil)

	appID := 1
	items, err := mockStorage.GetAllByAppID(appID)

	if err != nil {
		t.Fatalf("Se esperaba que no ocurriera un error, pero ocurrió: %v", err)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Los resultados obtenidos no coinciden con los valores esperados.")
	}

	mockStorage.AssertCalled(t, "GetAllByAppID", 1)

	mockStorage.AssertExpectations(t)
}
