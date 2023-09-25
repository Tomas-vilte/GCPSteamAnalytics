package api

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/controller"
	"github.com/gin-gonic/gin"
)

func MapRoutesCreateApiKey(r *gin.Engine) {
	r.POST("/createApiKey", controller.CreateKey)
}

func SetupRouteCreateKey(r *gin.Engine) {
	MapRoutesCreateApiKey(r)
}
