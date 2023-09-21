package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutesGetGames(r *gin.Engine, app controller.ProcessController) {
	r.POST("/processGames", app.Process)
}

func SetupRoutesGetGamesGCP(r *gin.Engine, gameController controller.ProcessController) {
	MapRoutesGetGames(r, gameController)
}
