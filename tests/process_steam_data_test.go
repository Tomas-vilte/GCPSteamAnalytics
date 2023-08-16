package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockHTTPClients struct {
	mock.Mock
}

func (m *MockHTTPClients) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestProcessSteamData(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	defer db.Close()

	// Crear un mock de HTTPClient
	mockHTTPClient := new(MockHTTPClient)

	// Configurar el comportamiento esperado del mock para la llamada HTTP
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"123": {"success": true, "data": {"steam_appid": 123, "type": "game"}}}`)),
	}
	mockHTTPClient.On("Do", mock.Anything).Return(resp, nil)

	// Configurar el comportamiento esperado del mock para la consulta SELECT
	expectedAppID := 123
	mockDB.ExpectQuery("SELECT last_appid FROM state_table where last_appid = ?").
		WithArgs(expectedAppID).
		WillReturnRows(sqlmock.NewRows([]string{"steam_appid", "type"}).AddRow(expectedAppID, "game"))

	// Configurar el comportamiento esperado del mock para la consulta UPDATE (SaveLastProcessedAppid)
	mockDB.ExpectExec("UPDATE state_table SET last_appid = ?").
		WithArgs(expectedAppID).
		WillReturnResult(sqlmock.NewResult(1, 1)) // 1 row affected

	// Crear una instancia de SteamAPI con los mocks de base de datos y HTTPClient
	steamAPI := &steamapi.SteamAPI{DB: db, Client: mockHTTPClient}

	// Llamar a la función que estamos probando
	appIDs := []int{123}
	limit := 10
	result, err := steamAPI.ProcessSteamData(appIDs, limit)

	// Realizar aserciones sobre los resultados
	assert.NoError(t, err)
	assert.Equal(t, len(appIDs), len(result))

	// Asegurarse de que las llamadas esperadas a los mocks se realizaron
	mockHTTPClient.AssertExpectations(t)
	mockDB.ExpectationsWereMet()
}
