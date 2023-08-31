package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, app controller.ProcessController) {
	r.POST("/processGames", app.Process)
}

func SetupRoutes(r *gin.Engine, app controller.ProcessController) {
	MapRoutes(r, app)
}
