package GCPSteamAnalytics

import (
	"github.com/Tomas-vilte/GCPSteamAnalytics/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateApiKey(w http.ResponseWriter, r *http.Request) {
	rGin := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	api.SetupRouteCreateKey(rGin)
	rGin.Use(api.RateLimitMiddleware())
	rGin.ServeHTTP(w, r)
}
