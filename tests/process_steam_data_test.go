package tests

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/models"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockSteamData es una implementación simulada de la interfaz SteamData
type MockSteamData struct {
	mock.Mock
}

func (m *MockSteamData) ProcessSteamData(appIDs []int, limit int) ([]*models.AppDetails, error) {
	args := m.Called(appIDs, limit)
	return args.Get(0).([]*models.AppDetails), args.Error(1)
}

func TestProcessSteamData(t *testing.T) {
	// Creamos una instancia simulada de SteamData
	mockData := &MockSteamData{}
	expectedDetails := []*models.AppDetails{
		{
			SteamAppid:  123,
			Description: "Example description 1",
			Type:        "game",
			Name:        "Game 1",
			Publishers:  []string{"Publisher A", "Publisher B"},
			Developers:  []string{"Developer X"},
			IsFree:      false,
		},
		{
			SteamAppid:  456,
			Description: "Example description 2",
			Type:        "software",
			Name:        "Software 1",
			Publishers:  []string{"Publisher C"},
			Developers:  []string{"Developer Y", "Developer Z"},
			IsFree:      true,
		},
		{
			SteamAppid:  789,
			Description: "Example description 2",
			Type:        "software",
			Name:        "Software 1",
			Publishers:  []string{"Publisher C"},
			Developers:  []string{"Developer Y", "Developer Z"},
			IsFree:      true,
		},
	}
	mockData.On("ProcessSteamData", mock.AnythingOfType("[]int"), mock.AnythingOfType("int")).Return(expectedDetails, nil)

	// Llamamos a la función que se está probando
	appIDs := []int{123, 456, 789}
	limit := 10
	processedData, err := mockData.ProcessSteamData(appIDs, limit)

	// Agrega las afirmaciones según el comportamiento esperado
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Assertion: Verifica que la longitud de los datos procesados coincida con los detalles esperados
	if len(processedData) != len(expectedDetails) {
		t.Errorf("Expected processedData to have %d entries, but got: %d", len(expectedDetails), len(processedData))
	}

	mockData.AssertExpectations(t)

	// Assertion: Verifica que el método ProcessSteamData se haya llamado con los argumentos correctos
	mockData.AssertCalled(t, "ProcessSteamData", appIDs, limit)

	// Afirmaciones adicionales basadas en detalles específicos de los datos procesados
	for i, expected := range expectedDetails {
		if processedData[i].Name != expected.Name {
			t.Errorf("Processed data name mismatch at index %d. Expected: %s, Actual: %s", i, expected.Name, processedData[i].Name)
		}
	}
}
