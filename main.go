package GCPSteamAnalytics

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func createApp() controller.ProcessController {
	storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(&http.Client{})
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewProcessController(sv)
}

func StartServerProcess() {
	r := gin.Default()
	app := createApp()
	api.SetupRoutesGetGamesGCP(r, app)

	r.Run("localhost:8080")
}

func init() {
	log.Printf("App started!")
	StartServerProcess()
}
