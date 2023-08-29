package tests

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/db/mocks"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"testing"
)

func TestConnect(t *testing.T) {
	// Utiliza el mock en las pruebas
	database := &mocks.MockDatabase{}

	err := database.Connect()
	if err != nil {
		t.Errorf("Se esperaba una conexión exitosa, pero se obtuvo el siguiente error: %v", err)
	}

	if !database.Connected {
		t.Errorf("Se esperaba una conexión exitosa, pero la base de datos no está conectada")
	}
}

func TestClose(t *testing.T) {
	// Utiliza el mock en las pruebas
	database := &mocks.MockDatabase{}

	// Simula la conexión
	database.Connect()

	err := database.Close()
	if err != nil {
		t.Errorf("Se esperaba que se cerrara la conexión sin errores, pero se obtuvo el siguiente error: %v", err)
	}

	if database.Connected {
		t.Errorf("Se esperaba que la base de datos estuviera desconectada, pero aún está conectada")
	}
}

func TestConnectWithError(t *testing.T) {
	// Utiliza el mock en las pruebas
	database := &mocks.MockDatabase{
		ShouldFail: true, // Simula un error de conexión
	}

	err := database.Connect()
	if err == nil {
		t.Error("Se esperaba un error de conexión, pero no se recibió ningún error")
	} else {
		t.Logf("Error recibido: %v", err)
	}

	if database.Connected {
		t.Errorf("Se esperaba que la conexión fallara, pero la base de datos está conectada")
	}
}

func TestInsertBatchData_Success(t *testing.T) {
	// Crea el mock de la base de datos
	mockDB := &mocks.MockDatabase{}

	// Simula una conexión exitosa
	mockDB.Connect()

	// Datos de prueba para insertar en lotes
	items := []entity.Item{
		{Appid: 1, Name: "Juego 1"},
		{Appid: 2, Name: "Juego 2"},
	}

	// Configura el mock para simular una inserción exitosa
	mockDB.ShouldInsert = true

	// Ejecuta la función que se va a probar (InsertBatchData) con el mock como base de datos
	err := mockDB.InsertBatchData(items)
	if err != nil {
		t.Errorf("Se esperaba una inserción exitosa, pero ocurrió un error: %v", err)
	}

	// Verifica que los items se hayan insertado correctamente en el mock
	insertedItems := mockDB.GetInsertedItems()
	if len(insertedItems) != len(items) {
		t.Errorf("Número incorrecto de items insertados. Se esperaba %d, se obtuvo %d", len(items), len(insertedItems))
	}
}
