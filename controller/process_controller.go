package controller

import (
	"context"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProcessController interface {
	Process(ctx *gin.Context)
}

func NewProcessController(sv *service.GameProcessor) ProcessController {
	return &processController{
		sv: sv,
	}
}

type processController struct {
	sv *service.GameProcessor
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
