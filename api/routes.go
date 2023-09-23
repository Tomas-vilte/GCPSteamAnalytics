package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, app controller.ProcessController, reviewCtrl controller.ReviewController, gameController controller.GameController) {
	r.POST("/processGames", app.Process)
	r.POST("/processReviews", reviewCtrl.ProcessReviews)
	r.GET("/getReviews", reviewCtrl.GetReviews)
	r.GET("/gameDetails", gameController.GetGameDetailsByID)
	r.GET("/games", gameController.GetGames)
}

func SetupRoutes(r *gin.Engine, app controller.ProcessController, reviewCtrl controller.ReviewController, gameController controller.GameController) {
	MapRoutes(r, app, reviewCtrl, gameController)
}
