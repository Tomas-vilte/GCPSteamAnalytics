package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, app controller.ProcessController, reviewCtrl controller.ReviewController) {
	r.POST("/processGames", app.Process)
	r.POST("/fetchReviews", reviewCtrl.FetchReviews)
}

func SetupRoutes(r *gin.Engine, app controller.ProcessController, reviewCtrl controller.ReviewController) {
	MapRoutes(r, app, reviewCtrl)
}
