package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ReviewController interface {
	ProcessReviews(ctx *gin.Context)
	GetReviews(ctx *gin.Context)
}

type reviewControllers struct {
	reviewAPI   service.ReviewsClient
	dbClient    persistence.StorageDB
	redisClient cache.RedisClient
}

func NewReviewController(api service.ReviewsClient, storage persistence.StorageDB, redisClient cache.RedisClient) ReviewController {
	return &reviewControllers{
		reviewAPI:   api,
		dbClient:    storage,
		redisClient: redisClient,
	}
}

func (rc *reviewControllers) ProcessReviews(ctx *gin.Context) {
	typeReview := ctx.DefaultQuery("review_type", "negative")
	appidStr := ctx.DefaultQuery("appid", "10")
	limit := ctx.DefaultQuery("limit", "10")
	appid, err := strconv.Atoi(appidStr)
	if err != nil {
		log.Printf("Error al convertir appid a int: %v", err)
		ctx.JSON(400, gin.H{
			"El valor de appid no es un número válido": err.Error(),
		})
		return
	}

	reviews, err := rc.reviewAPI.FetchReviews(appid, typeReview, limit)
	if err != nil {
		log.Printf("Error al obtener las revisiones para appID %d: %v", appid, err)
		ctx.JSON(500, gin.H{
			"Error al obtener las revisiones": err.Error(),
		})
		return
	}

	err = rc.dbClient.InsertReviews(appid, typeReview, reviews.Reviews)
	if err != nil {
		log.Printf("Error al insertar las revisiones en la base de datos: %v", err)
		ctx.JSON(500, gin.H{
			"Error al insertar las revisiones en la base de datos": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, reviews)
}

func (rc *reviewControllers) GetReviews(ctx *gin.Context) {
	appID, typeReview, limit, err := parseURLParams(ctx)
	if err != nil {
		log.Printf("Error al analizar los parámetros de la URL: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error al analizar los parámetros de la URL": err.Error(),
		})
		return
	}

	cacheKey := fmt.Sprintf("reviews:%d:%s:%d", appID, typeReview, limit)
	cachedReviews, err := rc.redisClient.GetReviewsInCache(cacheKey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"Error al obtener reseñas desde la caché": err.Error(),
		})
		return
	}

	if cachedReviews != nil {
		// Datos encontrados en caché, devolverlos directamente
		ctx.JSON(http.StatusOK, *cachedReviews)
		return
	}

	reviews, totalReviews, err := rc.dbClient.GetReviews(appID, typeReview, limit)
	if err != nil {
		log.Printf("Error al obtener reseñas desde la base de datos: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error al obtener reseñas desde la base de datos": err.Error(),
		})
		return
	}
	metadata := generateMetadataReview(totalReviews, limit, typeReview)
	response := steamapi.ReviewsResponse{
		Metadata: metadata,
		Reviews:  reviews,
	}

	// Almacenar los datos en Redis con un tiempo de expiración
	jsonData, err := json.Marshal(response)
	if err == nil {
		err := rc.redisClient.Set(cacheKey, string(jsonData))
		if err != nil {
			return
		}
	}

	ctx.JSON(200, response)
}

func parseURLParams(ctx *gin.Context) (int, string, int, error) {
	// Obtener los parámetros de la URL
	appidStr := ctx.DefaultQuery("appid", "10")
	typeReview := ctx.DefaultQuery("review_type", "positive")
	limitStr := ctx.DefaultQuery("limit", "10")

	// Convierte appid a int
	appid, err := strconv.Atoi(appidStr)
	if err != nil {
		return 0, "", 0, err
	}

	// Convierte limit a int
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, "", 0, err
	}

	return appid, typeReview, limit, nil
}

func generateMetadataReview(totalReview, size int, typeReview string) map[string]interface{} {
	metadata := map[string]interface{}{
		"size":         size,
		"total_review": totalReview,
		"type_review":  typeReview,
	}

	return metadata
}
