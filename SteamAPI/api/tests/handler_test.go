package tests

import (
	"github.com/stretchr/testify/assert"
	"steamAPI/api/handlers"
	"steamAPI/api/handlers/mocks"
	"testing"
)

// Implementa una estructura de prueba Mock para DataFetcher.
type MockDataFetcher struct{}

// GetData simula la obtención de datos y retorna datos simulados para tus pruebas.
func (m *MockDataFetcher) GetData() ([]handlers.Item, error) {
	mockItems := []handlers.Item{
		{Appid: 1, Name: "Item 1"},
		{Appid: 2, Name: "Item 2"},
	}
	return mockItems, nil
}

// GetDataWithError simula un error al obtener datos y retorna un error personalizado.
func (m *MockDataFetcher) GetDataWithError() ([]handlers.Item, error) {
	return nil, handlers.ErrDataFetch
}

// TestGetDataWithMock verifica que la función GetData utilice el MockDataFetcher para obtener datos simulados.
func TestGetDataWithMock(t *testing.T) {
	// Crea una instancia del mock.
	mockDataFetcher := &MockDataFetcher{}

	// Usa el mock en lugar de la implementación real.
	items, err := mockDataFetcher.GetData()

	// Realiza las aserciones de prueba con los datos simulados (items) y el error.
	if err != nil {
		t.Errorf("Se esperaba un error nulo, pero se obtuvo: %v", err)
	}

	// Asegúrate de que se obtuvieron los elementos simulados correctamente.
	expectedItems := []handlers.Item{
		{Appid: 1, Name: "Item 1"},
		{Appid: 2, Name: "Item 2"},
	}

	// Compara las listas de elementos obtenidas y esperadas.
	if len(items) != len(expectedItems) {
		t.Errorf("El número de elementos obtenidos no coincide con los esperados.")
	}

	for i := range items {
		if items[i].Appid != expectedItems[i].Appid || items[i].Name != expectedItems[i].Name {
			t.Errorf("El elemento obtenido no coincide con el esperado en la posición %d.", i)
		}
	}
}

// TestGetDataErrorHandling verifica que la función GetData maneje adecuadamente el error devuelto por GetDataWithError.
func TestGetDataErrorHandling(t *testing.T) {
	// Crea una instancia del mock.
	mockDataFetcher := &MockDataFetcher{}

	// Usa el mock para simular un error al obtener datos.
	_, err := mockDataFetcher.GetDataWithError()

	// Realiza aserciones para asegurarte de que el error se maneja adecuadamente.
	if err == nil {
		t.Errorf("Se esperaba un error, pero se obtuvo nulo.")
	} else {
		// Verificar el tipo de error.
		if err != handlers.ErrDataFetch {
			t.Errorf("Se esperaba un error de tipo handlers.ErrDataFetch.")
		}

		// Verificar el mensaje de error.
		expectedErrorMsg := "error al obtener datos"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Mensaje de error incorrecto. Se esperaba '%s', pero se obtuvo '%s'.", expectedErrorMsg, err.Error())
		}
	}
}

func TestInsertData_Success(t *testing.T) {
	// Crear el mock del DataFetcher
	mockDataFetcher := &mocks.MockDataFetcher{}

	// Crear el mock de la base de datos
	mockDB := &mocks.MockDatabase{}

	// Simular una conexión exitosa en el mock de la base de datos
	mockDB.Connected = true

	// Datos de prueba que devolverá el mock del DataFetcher
	items := []handlers.Item{
		{Appid: 1, Name: "Juego 1"},
		{Appid: 2, Name: "Juego 2"},
	}

	// Configurar el mock del DataFetcher para devolver los items de prueba
	mockDataFetcher.On("GetData").Return(items, nil)

	// Configurar el mock de la base de datos para simular una inserción exitosa
	mockDB.ShouldInsert = true

	// Ejecuta la función que se va a probar (InsertData) con los mocks
	err := mockDB.InsertData(mockDataFetcher)
	assert.NoError(t, err, "Se esperaba una inserción exitosa")

	// Verificar que los items se hayan insertado correctamente en el mock de la base de datos
	insertedItems := mockDB.GetInsertedItems()
	assert.Equal(t, len(items), len(insertedItems), "Número incorrecto de items insertados")

}
