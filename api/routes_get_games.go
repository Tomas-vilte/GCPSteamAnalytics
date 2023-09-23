package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutesGetGamesFromDB(r *gin.Engine, app controller.GameController) {
	r.GET("/gameDetails", app.GetGameDetailsByID)
	r.GET("/games", app.GetGames)
}

func SetupRoutesGetGamesFromDB(r *gin.Engine, app controller.GameController) {
	MapRoutesGetGamesFromDB(r, app)
}
