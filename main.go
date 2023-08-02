package GCPSteamAnalytics

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Tomas-vilte/GCPSteamAnalytics/functionGCP"
)

func init() {
	functions.HTTP("MyCloudFunction", functionGCP.MyCloudFunction)
}
