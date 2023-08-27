package tests

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	_ "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveToCSV(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "csv_test")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFilePath := filepath.Join(tempDir, "test.csv")

	testData := []model.AppDetails{
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
	}

	steamAPI := &steamapi.SteamAPI{}

	err = steamAPI.SaveToCSV(testData, tempFilePath)
	if err != nil {
		t.Fatalf("Error calling SaveToCSV: %v", err)
	}

	fileContent, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	expectedContent := `SteamAppid,Description,Type,Name,Publishers,Developers,isFree,Windows,Mac,Linux,Date,ComingSoon,Currency,DiscountPercent,InitialFormatted,FinalFormatted
123,Example description 1,game,Game 1,"Publisher A, Publisher B",Developer X,false,false,true,false,2023-08-15,false,USD,25,19.99,14.99 USD
456,Example description 2,software,Software 1,Publisher C,"Developer Y, Developer Z",true,true,true,true,2023-09-01,true,EUR,0,0,Free`

	if string(fileContent) == expectedContent {
		t.Errorf("El contenido del archivo no coincide con el valor esperado.\nExpected:\n%s\nActual:\n%s", expectedContent, fileContent)
	}
}
