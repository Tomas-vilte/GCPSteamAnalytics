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

func createGameDetails() controller.GameController {
	steamClient := service.NewSteamClient(&http.Client{})
	redisClient := cache.NewRedisCacheClient("localhost:6379", 1)
	storage := persistence.NewStorage()
	sv := service.NewGameProcessor(storage, steamClient)
	return controller.NewGameController(steamClient, redisClient, storage, *sv)

}

func GetGames(w http.ResponseWriter, r *http.Request) {
	rGIn := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app := createGameDetails()
	api.SetupRoutesGetGamesFromDB(rGIn, app)
	rGIn.ServeHTTP(w, r)
}
