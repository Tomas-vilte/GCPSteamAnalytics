package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type GameController interface {
	GetGameDetails(ctx *gin.Context)
}

type gameController struct {
	steamService service.SteamClient
	redisClient  cache.RedisClient
	dbClient     persistence.StorageDB
}

func NewGameController(steamService service.SteamClient, redisClient cache.RedisClient, db persistence.StorageDB) GameController {
	return &gameController{
		steamService: steamService,
		dbClient:     db,
		redisClient:  redisClient,
	}
}

func (gc *gameController) GetGameDetails(ctx *gin.Context) {
	gameID := ctx.Param("id")
	gameint, _ := strconv.Atoi(gameID)

	// Consultar Redis para ver si los detalles del juego están en caché.
	cachedDetails, err := gc.redisClient.Get(gameID)
	if err != nil {
		if err != redis.Nil {
			// Ocurrió un error diferente al intentar obtener datos de Redis.
			ctx.JSON(500, gin.H{
				"error": fmt.Sprintf("Error al obtener detalles del juego desde Redis: %v", err),
			})
			return
		}

		// Si el juego no esta en la cache, lo buscamos en la bd
		dbDetails, err := gc.dbClient.GetGameDetails(10)
		fmt.Println(dbDetails)
		if err != nil {
			if err == sql.ErrNoRows {
				// Si no esta en la bd, hacemos un api call a la api de steam
				apiDetails, err := gc.steamService.GetAppDetails(gameint)
				fmt.Println("SEXOOOOOOOOOOOOOOOOOO", apiDetails)
				if err != nil {
					// Ocurrió un error al intentar obtener datos de la API.
					ctx.JSON(500, gin.H{
						"error": fmt.Sprintf("Error al obtener detalles del juego desde la API: %v", err),
					})
					return
				}

				// Guardar los detalles obtenidos de la API en caché.
				err = gc.redisClient.Set(gameID, string(apiDetails))
				if err != nil {
					// Ocurrió un error al intentar guardar en caché los detalles de la API.
					ctx.JSON(500, gin.H{
						"error": fmt.Sprintf("Error al guardar detalles del juego en caché: %v", err),
					})
					return
				}

				// Responder con los detalles obtenidos de la API.
				ctx.JSON(200, apiDetails)
				return
			}

			// Ocurrió un error diferente al intentar obtener datos de la base de datos.
			ctx.JSON(500, gin.H{
				"error": fmt.Sprintf("Error al obtener detalles del juego desde la base de datos: %v", err),
			})
			return
		}
		jsonData, err := json.Marshal(dbDetails)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": fmt.Sprintf("Error al serializar detalles del juego: %v", err),
			})
			return
		}
		// Los detalles del juego se encontraron en la base de datos.
		// Guardarlos en caché para futuras consultas.
		err = gc.redisClient.Set(gameID, string(jsonData))
		if err != nil {
			// Ocurrió un error al intentar guardar en caché los detalles de la base de datos.
			ctx.JSON(500, gin.H{
				"error": fmt.Sprintf("Error al guardar detalles del juego en caché: %v", err),
			})
			return
		}

		// Responder con los detalles obtenidos de la base de datos.
		ctx.JSON(200, dbDetails)
		return
	}

	// Los detalles del juego se encontraron en caché.
	ctx.JSON(200, cachedDetails)
}
