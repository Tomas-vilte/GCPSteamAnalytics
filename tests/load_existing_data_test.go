package tests

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func createTempFile(content string) string {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	_, err = tempFile.WriteString(content)
	if err != nil {
		panic(err)
	}

	return tempFile.Name()
}

func TestLoadExistingData_FileNotFound(t *testing.T) {
	existingData, err := utils.LoadExistingData("nonexistent.csv")
	assert.NoError(t, err)
	assert.Empty(t, existingData)
}

func TestLoadExistingData_ValidFile(t *testing.T) {
	fileContent := "appID\n10\n20\n30\n"
	filePath := createTempFile(fileContent)
	defer os.Remove(filePath)

	existingData, err := utils.LoadExistingData(filePath)
	assert.NoError(t, err)
	assert.Len(t, existingData, 3)
	assert.True(t, existingData[10])
	assert.True(t, existingData[20])
	assert.True(t, existingData[30])
}

func TestLoadExistingData_InvalidData(t *testing.T) {
	fileContent := "appID\n10\ninvalid\n20\n"
	filePath := createTempFile(fileContent)
	defer os.Remove(filePath)

	existingData, err := utils.LoadExistingData(filePath)
	assert.NoError(t, err)
	assert.Len(t, existingData, 2)
	assert.True(t, existingData[10])
	assert.True(t, existingData[20])
}
