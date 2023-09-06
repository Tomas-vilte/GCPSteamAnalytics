package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	app := createApp()
	reviewCtrl := createReviewController()
	getGame := createGameDetails()
	SetupRoutes(r, app, reviewCtrl, getGame)

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

func createGameDetails() controller.GameController {
	steamClient := service.NewSteamClient(&http.Client{})
	redisClient := cache.NewRedisCacheClient("localhost:6379", 1)
	storage := persistence.NewStorage()
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewGameController(steamClient, redisClient, storage, *sv)

}
