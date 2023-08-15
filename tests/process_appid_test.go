package tests

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

// Define un mock para el cliente HTTP
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestProcessAppID(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/steamAnalytics")
	if err != nil {
		log.Printf("Error al abrir la base de datos: %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Hubo un error al conectarse a la base de datos: %v", err)
		defer db.Close()
		return
	}
	// Configura el cliente HTTP mock
	mockClient := new(MockHTTPClient)

	// Crea una instancia de SteamAPI con el cliente mock
	api := &steamapi.SteamAPI{
		DB:     db,
		Client: mockClient,
	}

	// Define la respuesta esperada del cliente mock
	expectedResponse := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"10": {"success": true, "data": {"steam_appid": 10, "type": "game"}}}`)),
	}

	// Configura el comportamiento del cliente mock
	mockClient.On("Do", mock.Anything).Return(expectedResponse, nil)

	// Llama a la función que deseas probar
	appDetails, err := api.ProcessAppID(10)

	// Verifica que no haya error
	assert.NoError(t, err)

	// Verifica que los detalles de la aplicación sean los esperados
	assert.NotNil(t, appDetails)
	assert.Equal(t, int64(10), appDetails.SteamAppid)
	assert.Equal(t, "game", appDetails.Type)

	// Verifica el comportamiento del cliente mock
	mockClient.AssertExpectations(t)
}

func TestProcessAppID_UpdateLastProcessedAppID(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockClient := new(MockHTTPClient)
	api := &steamapi.SteamAPI{
		DB:     db,
		Client: mockClient,
	}

	// Configura el comportamiento del mock de la base de datos para que no retorne errores y no se espere ninguna actualización
	mockDB.ExpectExec("UPDATE state_table SET last_appid = ?").
		WithArgs(10).
		WillReturnResult(sqlmock.NewResult(0, 0)) // Simula que no se actualizó nada

	expectedResponse := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"10": {"success": true, "data": {"steam_appid": 10, "type": "game"}}}`)),
	}
	mockClient.On("Do", mock.Anything).Return(expectedResponse, nil)

	// Llama a la función que deseas probar
	appDetails, err := api.ProcessAppID(10)

	// Verifica que no haya error
	assert.NoError(t, err)
	assert.NotNil(t, appDetails)

	// Asegúrate de que todas las expectativas del mock se cumplan
	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}