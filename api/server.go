package api

import (
	"net/http"

	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	app := createApp()
	reviewCtrl := createReviewController()
	getGame := createGameDetails()

	r.Use(RateLimitMiddleware())

	SetupRoutes(r, app, reviewCtrl, getGame)
	r.Run("localhost:8081")
}

func createApp() controller.ProcessController {
	storage := persistence.NewStorage()
	steamClient := service.NewSteamClient(&http.Client{})
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewProcessController(sv)
}

func createReviewController() controller.ReviewController {
	sv := service.NewSteamReviewAPI(&http.Client{})
	storage := persistence.NewStorage()
	redisClient := cache.NewRedisCacheClient("localhost:6379", 1)
	return controller.NewReviewController(sv, storage, redisClient)
}

func createGameDetails() controller.GameController {
	steamClient := service.NewSteamClient(&http.Client{})
	redisClient := cache.NewRedisCacheClient("localhost:6379", 1)
	storage := persistence.NewStorage()
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewGameController(steamClient, redisClient, storage, *sv)

}
