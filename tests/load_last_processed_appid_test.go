package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadLastProcessedAppid(t *testing.T) {
	// Crear una base de datos simulada y un mock para las expectativas
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Definir el valor esperado
	expectedAppID := 10

	// Configurar la expectativa para la consulta SQL
	mock.ExpectQuery("SELECT last_appid FROM state_table").
		WillReturnRows(sqlmock.NewRows([]string{"last_appid"}).AddRow(expectedAppID))

	// Crear una instancia de SteamAPI para probar
	api := &steamapi.SteamAPI{DB: db}

	// Llamar a la función bajo prueba
	appID, err := api.LoadLastProcessedAppid()

	// Verificar los resultados y las expectativas del mock
	assert.NoError(t, err)
	assert.Equal(t, expectedAppID, appID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
