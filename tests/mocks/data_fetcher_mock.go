package mocks

import (
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/handlers"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/stretchr/testify/mock"
)

// InsertBatchData Implementar la función InsertBatchData para el mock
func (m *MockDatabase) InsertBatchData(items []entity.Item) error {
	if m.ShouldFail {
		return errors.New("error de inserción simulado")
	}

	// Simular una inserción exitosa o fallida según el valor de ShouldInsert
	if m.ShouldInsert {
		m.InsertedItems = append(m.InsertedItems, items...)
	} else {
		m.InsertedItems = []entity.Item{} // Simular una inserción fallida eliminando los items insertados
	}

	return nil
}

type MockDataFetcher struct {
	mock.Mock
}

func (m *MockDataFetcher) GetData() ([]entity.Item, error) {
	// Utiliza la biblioteca de mocks para configurar el comportamiento del método GetData
	args := m.Called()
	return args.Get(0).([]entity.Item), args.Error(1)
}

// Función de ayuda para obtener los items insertados en el mock
func (m *MockDatabase) GetInsertedItems() []entity.Item {
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
