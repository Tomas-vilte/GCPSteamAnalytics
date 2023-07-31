package mocks

import (
	"errors"

	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
)

type MockDatabase struct {
	Connected     bool
	ShouldFail    bool            // Nuevo campo para indicar si debe simular un error
	InsertedItems []handlers.Item // Nuevo campo para almacenar los items insertados en el mock
	ShouldInsert  bool            // Nuevo campo para indicar si debe simular una inserción exitosa o fallida
}

func (m *MockDatabase) Connect() error {
	// Simula una conexión exitosa a menos que ShouldFail sea verdadero
	if m.ShouldFail {
		return errors.New("error de conexión simulado")
	}

	m.Connected = true
	return nil
}

func (m *MockDatabase) Close() error {
	// Simular el cierre de la conexión
	m.Connected = false
	return nil
}

// InsertBatchData Implementar la función InsertBatchData para el mock
func (m *MockDatabase) InsertBatchData(items []handlers.Item) error {
	if m.ShouldFail {
		return errors.New("error de inserción simulado")
	}

	// Simular una inserción exitosa o fallida según el valor de ShouldInsert
	if m.ShouldInsert {
		m.InsertedItems = append(m.InsertedItems, items...)
	} else {
		m.InsertedItems = []handlers.Item{} // Simular una inserción fallida eliminando los items insertados
	}

	return nil
}

// GetInsertedItems Función de ayuda para obtener los items insertados en el mock
func (m *MockDatabase) GetInsertedItems() []handlers.Item {
	return m.InsertedItems
}
