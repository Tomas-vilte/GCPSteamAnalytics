package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
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

func ProcessGameDetailsResponse(responseData []byte) ([]steamapi.AppDetails, error) {
	// Define una estructura que coincida con la respuesta de la API Steam.
	var response map[string]steamapi.AppDetailsResponse

	// Decodifica la respuesta JSON en la estructura.
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, err
	}

	// Inicializa un slice de AppDetails para almacenar los detalles del juego.
	var gameDetails []steamapi.AppDetails

	// Itera a través de los datos de la respuesta y agrega cada juego a gameDetails.
	for _, appDetailsResponse := range response {
		gameDetails = append(gameDetails, appDetailsResponse.Data)
	}

	return gameDetails, nil
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
				if err == sql.ErrNoRows {
					// El juego no está en la base de datos, lo buscamos en la API de Steam.
					apiDetails, err := gc.steamService.GetAppDetails(gameint)
					if err != nil {
						ctx.JSON(500, gin.H{
							"Error al obtener detalles del juego desde la API": err.Error(),
						})
						return
					}

					// Guardar los detalles obtenidos de la API en caché.
					apiDetailsJSON, err := json.Marshal(apiDetails)
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
					ctx.JSON(200, apiDetails)
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

//func (gc *gameController) GetGameDetails(ctx *gin.Context) {
//	gameID := ctx.Param("appid")
//	gameint, _ := strconv.Atoi(gameID)
//
//	// Consultar Redis para ver si los detalles del juego están en caché.
//	cachedDetails, err := gc.redisClient.Get(gameID)
//	if err != nil {
//		if err != redis.Nil {
//			// Ocurrió un error diferente al intentar obtener datos de Redis.
//			ctx.JSON(500, gin.H{
//				"Error al obtener detalles del juego desde Redis:": err.Error(),
//			})
//			return
//		}
//
//		// Si el juego no está en la caché, intentamos obtenerlo de la BD.
//		dbDetails, err := gc.dbClient.GetGameDetails(gameint)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				// Si no está en la BD, hacemos una llamada a la API de Steam.
//				apiDetails, err := gc.steamService.GetAppDetails(gameint)
//
//				if err != nil {
//					// Ocurrió un error al intentar obtener datos de la API.
//					ctx.JSON(500, gin.H{
//						"Error al obtener detalles del juego desde la API": err.Error(),
//					})
//					return
//				}
//
//				// Guardar los detalles obtenidos de la API en caché.
//				err = gc.redisClient.Set(gameID, string(apiDetails))
//				if err != nil {
//					// Ocurrió un error al intentar guardar en caché los detalles de la API.
//					ctx.JSON(500, gin.H{
//						"Error al guardar detalles del juego en caché:": err.Error(),
//					})
//					return
//				}
//				var apiDetails1 []model.AppDetails
//				if err := json.Unmarshal(apiDetails, &apiDetails); err != nil {
//					ctx.JSON(500, gin.H{
//						"Error al deserializar los detalles del juego desde la API": err.Error(),
//					})
//					return
//				}
//
//				// Guardar los detalles obtenidos de la API en la BD.
//				err = gc.dbClient.SaveGameDetails(apiDetails1)
//				if err != nil {
//					// Ocurrió un error al intentar guardar los detalles del juego en la BD.
//					ctx.JSON(500, gin.H{
//						"Error al guardar detalles del juego en la BD": err.Error(),
//					})
//					return
//				}
//
//				//// Actualizar el estado del juego en la BD como "procesado".
//				//err = gc.dbClient.UpdateGameStatus(gameint, entity.PROCESSED)
//				//if err != nil {
//				//	// Ocurrió un error al intentar actualizar el estado del juego.
//				//	ctx.JSON(500, gin.H{
//				//		"Error al actualizar el estado del juego en la BD": err.Error(),
//				//	})
//				//	return
//				//}
//
//				// Responder con los detalles obtenidos de la API.
//				ctx.JSON(200, apiDetails)
//				return
//			}
//
//			// Ocurrió un error diferente al intentar obtener datos de la base de datos.
//			ctx.JSON(500, gin.H{
//				"Error al obtener detalles del juego desde la base de datos:": err.Error(),
//			})
//			return
//		}
//
//		jsonData, err := json.Marshal(dbDetails)
//		if err != nil {
//			ctx.JSON(500, gin.H{
//				"Error al serializar detalles del juego:": err.Error(),
//			})
//			return
//		}
//
//		// Si el juego está en la BD, guardarlo en caché para futuras consultas.
//		err = gc.redisClient.Set(gameID, string(jsonData))
//		if err != nil {
//			// Ocurrió un error al intentar guardar en caché los detalles de la base de datos.
//			ctx.JSON(500, gin.H{
//				"Error al guardar detalles del juego en caché": err.Error(),
//			})
//			return
//		}
//
//		// Responder con los detalles obtenidos de la base de datos.
//		ctx.JSON(200, dbDetails)
//		return
//	}
//	// Los detalles del juego se encontraron en caché.
//	ctx.JSON(200, cachedDetails)
//}
