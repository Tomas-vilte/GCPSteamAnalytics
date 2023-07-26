package tests

import (
	"steamAPI/api/db/mocks"
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
