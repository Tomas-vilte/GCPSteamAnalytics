package controller

import (
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ReviewController interface {
	FetchReviews(ctx *gin.Context)
}

type reviewController struct {
	reviewAPI service.ReviewsClient
}

func NewReviewController(api service.ReviewsClient) ReviewController {
	return &reviewController{
		reviewAPI: api,
	}
}

func (rc *reviewController) FetchReviews(ctx *gin.Context) {
	typeReview := ctx.DefaultQuery("typeReview", "")
	appidStr := ctx.DefaultQuery("appid", "")
	appid, err := strconv.Atoi(appidStr)
	if err != nil {
		log.Printf("Error al convertir appid a int: %v", err)
		ctx.JSON(400, gin.H{
			"error": fmt.Sprintf("El valor de appid no es un número válido: %v", err),
		})
		return
	}

	reviews, err := rc.reviewAPI.FetchReviews(appid, typeReview)
	if err != nil {
		log.Printf("Error al obtener las revisiones para appID %d: %v", appid, err)
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("Error al obtener las revisiones para appID %d: %v", appid, err),
		})
		return
	}
	ctx.JSON(http.StatusOK, reviews)
}