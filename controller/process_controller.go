package controller

import (
	"context"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProcessController interface {
	Process(ctx *gin.Context)
	GetReviews(ctx *gin.Context)
}

func NewProcessController(sv *service.GameProcessor, reviewAPI *steamapi.SteamReviewAPI) ProcessController {
	return &processController{
		sv:        sv,
		reviewAPI: reviewAPI,
	}
}

type processController struct {
	sv        *service.GameProcessor
	reviewAPI *steamapi.SteamReviewAPI
}

func (p *processController) GetReviews(ctx *gin.Context) {
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

	reviews, err := p.reviewAPI.GetReviews(appid, typeReview)
	if err != nil {
		log.Printf("Error al obtener las revisiones para appID %d: %v", appid, err)
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("Error al obtener las revisiones para appID %d: %v", appid, err),
		})
		return
	}
	ctx.JSON(http.StatusOK, reviews)
}

func (p *processController) Process(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Valor límite no válido",
		})
		return
	}

	err = p.sv.RunProcessData(context.Background(), limit)
	if err != nil {
		log.Printf("Hubo un error: %v", err)
		ctx.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d registros de datos procesados", limit),
	})
}
