package GCPSteamAnalytics

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createApp() controller.ProcessController {
	storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(&http.Client{})
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewProcessController(sv)
}

func MyCloudFunction(w http.ResponseWriter, r *http.Request) {
	rGin := gin.Default()
	app := createApp()
	api.SetupRoutesGetGamesGCP(rGin, app)
	rGin.ServeHTTP(w, r)
}
