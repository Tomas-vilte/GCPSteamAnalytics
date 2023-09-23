package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutesGetReviews(r *gin.Engine, app controller.ReviewController) {
	r.POST("/processReviews", app.ProcessReviews)
	r.GET("/getReviews", app.GetReviews)
}

func SetupRoutesReviews(r *gin.Engine, app controller.ReviewController) {
	MapRoutesGetReviews(r, app)
}
