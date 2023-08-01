package GCPSteamAnalytics

import (
	"fmt"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Tomas-vilte/GCPSteamAnalytics/functionGCP"
)

func init() {
	functions.HTTP("ProcessSteamDataAndSaveToStorage", functionGCP.ProcessSteamDataAndSaveToStorage)
	fmt.Println("Function initialized")
}
