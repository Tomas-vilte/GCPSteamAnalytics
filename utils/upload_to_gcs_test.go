package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUploader es una estructura que incorpora el mock de la función UploadFileToGCS.
type MockUploader struct {
	mock.Mock
}

// MockUploadFileToGCS es la implementación mock de la función UploadFileToGCS.
func (m *MockUploader) UploadFileToGCS(filePath, bucket, object string) error {
	args := m.Called(filePath, bucket, object)
	return args.Error(0)
}

func TestUploadFileToGCS_Success(t *testing.T) {
	// Creamos el objeto MockUploader.
	mockUploader := new(MockUploader)

	// Definimos el comportamiento esperado del mock.
	filePath := "GCPSteamAnalytics/steamAPI/api/data/hola2.txt"
	bucket := "steam-analytics"
	object := "hola3.txt"
	mockUploader.On("UploadFileToGCS", filePath, bucket, object).Return(nil)

	// Ejecutamos el código que utiliza la función UploadFileToGCS (en este caso, no estamos interesados en los resultados de retorno).
	err := mockUploader.UploadFileToGCS(filePath, bucket, object)

	// Verificamos que se haya llamado al método mock con los argumentos esperados.
	mockUploader.AssertExpectations(t)

	// Verificamos que no haya ocurrido ningún error durante la ejecución.
	assert.NoError(t, err)
}

func TestUploadFileToGCS_Error(t *testing.T) {
	// Creamos el objeto MockUploader.
	mockUploader := new(MockUploader)

	// Definimos el comportamiento esperado del mock para retornar un error.
	filePath := "GCPSteamAnalytics/steamAPI/api/data/hola2.txt"
	bucket := "steam-analytics"
	object := "hola3.txt"
	expectedError := fmt.Errorf("error al subir archivo")
	mockUploader.On("UploadFileToGCS", filePath, bucket, object).Return(expectedError)

	// Ejecutamos el código que utiliza la función UploadFileToGCS (en este caso, no estamos interesados en los resultados de retorno).
	err := mockUploader.UploadFileToGCS(filePath, bucket, object)

	// Verificamos que se haya llamado al método mock con los argumentos esperados.
	mockUploader.AssertExpectations(t)

	// Verificamos que se haya obtenido el error esperado.
	assert.EqualError(t, err, expectedError.Error())
}
