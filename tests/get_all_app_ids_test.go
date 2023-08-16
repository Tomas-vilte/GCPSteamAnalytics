package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllAppIDs(t *testing.T) {
	// Crear una base de datos simulada y un mock para las expectativas
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Definir valores esperados
	lastProcessedAppID := 5
	expectedAppIDs := []int{6, 7, 8}

	// Configurar la expectativa para la consulta SQL
	mock.ExpectQuery("SELECT appid FROM games WHERE appid > ?").
		WithArgs(lastProcessedAppID).
		WillReturnRows(sqlmock.NewRows([]string{"appid"}).
			AddRow(6).
			AddRow(7).
			AddRow(8))

	// Crear una instancia de SteamAPI para probar
	api := &steamapi.SteamAPI{DB: db}

	// Llamar a la función bajo prueba
	appIDs, err := api.GetAllAppIDs(lastProcessedAppID)

	// Verificar los resultados y las expectativas del mock
	assert.NoError(t, err)
	assert.Equal(t, expectedAppIDs, appIDs)
	assert.NoError(t, mock.ExpectationsWereMet())
}
