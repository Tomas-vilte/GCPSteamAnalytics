package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveLastProcessedAppid(t *testing.T) {
	// Crear una base de datos simulada y un mock para las expectativas
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Definir el valor esperado para la actualización
	lastProcessedAppID := 15

	// Configurar la expectativa para la consulta SQL
	mock.ExpectExec("UPDATE state_table SET last_appid = ?").
		WithArgs(lastProcessedAppID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Crear una instancia de SteamAPI para probar
	api := &steamapi.SteamAPI{DB: db}

	// Llamar a la función bajo prueba
	err = api.SaveLastProcessedAppid(lastProcessedAppID)

	// Verificar los resultados y las expectativas del mock
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
