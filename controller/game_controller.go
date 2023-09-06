package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

type GameController interface {
	GetGameDetails(ctx *gin.Context)
}

type gameController struct {
	steamService service.SteamClient
	redisClient  cache.RedisClient
	gameProcesor service.GameProcessor
	dbClient     persistence.StorageDB
}

func NewGameController(steamService service.SteamClient, redisClient cache.RedisClient, db persistence.StorageDB, gameProcesor service.GameProcessor) GameController {
	return &gameController{
		steamService: steamService,
		dbClient:     db,
		gameProcesor: gameProcesor,
		redisClient:  redisClient,
	}
}

func (gc *gameController) GetGameDetails(ctx *gin.Context) {
	gameID := ctx.Param("appid")
	gameint, _ := strconv.Atoi(gameID)

	// Consultar Redis para ver si los detalles del juego están en caché.
	cachedDetails, err := gc.redisClient.Get(gameID)
	if err != nil {
		if err == redis.Nil {
			// El juego no está en caché, intentamos obtenerlo de la base de datos.
			dbDetails, err := gc.dbClient.GetGameDetails(gameint)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					// El juego no está en la base de datos, lo buscamos en la API de Steam.
					apiDetails, err := gc.steamService.GetAppDetails(gameint)
					if err != nil {
						ctx.JSON(500, gin.H{
							"Error al obtener detalles del juego desde la API": err.Error(),
						})
						return
					}
					apiDetailsSlice := [][]byte{apiDetails}
					games, err := gc.dbClient.GetAllByAppID(gameint)
					responseData, err := gc.gameProcesor.ProcessResponse(apiDetailsSlice, games)
					if err != nil {
						log.Printf("error: %v", err)
					}

					err = gc.dbClient.SaveGameDetails(responseData)
					if err != nil {
						log.Printf("error: %v", err)
					}

					// Guardar los detalles obtenidos de la API en caché.
					apiDetailsJSON, err := json.Marshal(responseData)
					if err != nil {
						ctx.JSON(500, gin.H{
							"Error al serializar detalles del juego desde la API": err.Error(),
						})
						return
					}
					err = gc.redisClient.Set(gameID, string(apiDetailsJSON))
					if err != nil {
						ctx.JSON(500, gin.H{
							"Error al guardar detalles del juego en caché desde la API": err.Error(),
						})
						return
					}

					// Responder con los detalles obtenidos de la API.
					ctx.JSON(200, responseData)
					return
				}

				// Ocurrió un error diferente al intentar obtener datos de la base de datos.
				ctx.JSON(500, gin.H{
					"Error al obtener detalles del juego desde la base de datos": err.Error(),
				})
				return
			}

			// Los detalles del juego se encontraron en la base de datos.
			// Guardarlos en caché para futuras consultas.
			dbDetailsJSON, err := json.Marshal(dbDetails)
			if err != nil {
				ctx.JSON(500, gin.H{
					"Error al serializar detalles del juego desde la base de datos": err.Error(),
				})
				return
			}
			err = gc.redisClient.Set(gameID, string(dbDetailsJSON))
			if err != nil {
				ctx.JSON(500, gin.H{
					"Error al guardar detalles del juego en caché desde la base de datos": err.Error(),
				})
				return
			}

			// Responder con los detalles obtenidos de la base de datos.
			ctx.JSON(200, dbDetails)
			return
		}

		// Ocurrió un error diferente al intentar obtener datos de Redis.
		ctx.JSON(500, gin.H{
			"Error al obtener detalles del juego desde Redis": err.Error(),
		})
		return
	}

	// Los detalles del juego se encontraron en caché.
	ctx.JSON(200, cachedDetails)
}

func (gc *gameController) getCachedGameDetails(gameID int) (interface{}, error) {
	cachedDetails, err := gc.redisClient.Get(strconv.Itoa(gameID))
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		return nil, nil
	}
	return cachedDetails, nil
}
