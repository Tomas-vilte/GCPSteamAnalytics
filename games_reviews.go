package GCPSteamAnalytics

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createReviewController() controller.ReviewController {
	sv := service.NewSteamReviewAPI(&http.Client{})
	storage := persistence.NewStorage()
	redisClient := cache.NewRedisCacheClient("localhost:6379", 1)
	return controller.NewReviewController(sv, storage, redisClient)
}

func GamesReviews(w http.ResponseWriter, r *http.Request) {
	rGin := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app := createReviewController()
	api.SetupRoutesReviews(rGin, app)
	rGin.Use(api.RateLimitMiddleware())
	rGin.ServeHTTP(w, r)
}
