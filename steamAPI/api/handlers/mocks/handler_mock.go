package mocks

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"steamAPI/api/handlers"
)

type MockDatabase struct {
	Connected    bool
	ShouldFail   bool // Nuevo campo para indicar si debe simular un error de conexión
	ShouldInsert bool // Nuevo campo para indicar si debe simular una inserción exitosa o fallida

	// Nuevo campo para almacenar los items insertados en el mock
	InsertedItems []handlers.Item
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

type MockDataFetcher struct {
	mock.Mock
}

func (m *MockDataFetcher) GetData() ([]handlers.Item, error) {
	// Utiliza la biblioteca de mocks para configurar el comportamiento del método GetData
	args := m.Called()
	return args.Get(0).([]handlers.Item), args.Error(1)
}

// Función de ayuda para obtener los items insertados en el mock
func (m *MockDatabase) GetInsertedItems() []handlers.Item {
	return m.InsertedItems
}

// Implementa la función InsertData para que sea utilizada en la prueba
func (m *MockDatabase) InsertData(dataFetcher handlers.DataFetcher) error {
	// Obtén los datos del DataFetcher
	items, err := dataFetcher.GetData()
	if err != nil {
		return err
	}

	// Inserta los datos en lotes en la base de datos
	err = m.InsertBatchData(items)
	if err != nil {
		return err
	}

	return nil
}
