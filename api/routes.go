package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, app controller.ProcessController, reviewCtrl controller.ReviewController, gameController controller.GameController) {
	r.POST("/processGames", app.Process)
	r.POST("/fetchReviews", reviewCtrl.FetchReviews)
	r.GET("/gamedetails/:appid", gameController.GetGameDetailsByID)
}

func SetupRoutes(r *gin.Engine, app controller.ProcessController, reviewCtrl controller.ReviewController, gameController controller.GameController) {
	MapRoutes(r, app, reviewCtrl, gameController)
}
