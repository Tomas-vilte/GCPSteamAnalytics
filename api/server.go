package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer() {
	r := gin.Default()
	app := createApp()
	reviewCtrl := createReviewController()
	SetupRoutes(r, app, reviewCtrl)

	r.Run("localhost:8080")
}

func createApp() controller.ProcessController {
	storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(&http.Client{})
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewProcessController(sv)
}

func createReviewController() controller.ReviewController {
	sv := service.NewSteamReviewAPI(&http.Client{})
	return controller.NewReviewController(sv)
}
