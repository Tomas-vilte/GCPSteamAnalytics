package GCPSteamAnalytics

import (
	"net/http"

	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
)

func createApp() controller.ProcessController {
	storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(&http.Client{})
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewProcessController(sv)
}

func ProcessGames(w http.ResponseWriter, r *http.Request) {
	rGin := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app := createApp()
	api.SetupRoutesGetGamesGCP(rGin, app)
	rGin.ServeHTTP(w, r)
}
