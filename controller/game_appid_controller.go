package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

type GameController interface {
	GetGameDetailsByID(ctx *gin.Context)
	GetGames(c *gin.Context)
}

type GameControllers struct {
	steamService  service.SteamClient
	redisClient   cache.RedisClient
	gameProcessor service.GameProcessor
	dbClient      persistence.StorageDB
}

func NewGameController(steamService service.SteamClient, redisClient cache.RedisClient, db persistence.StorageDB, gameProcessor service.GameProcessor) GameController {
	return &GameControllers{
		steamService:  steamService,
		dbClient:      db,
		gameProcessor: gameProcessor,
		redisClient:   redisClient,
	}
}

func (gc *GameControllers) GetGameDetailsByID(ctx *gin.Context) {
	gameID := ctx.Param("appid")
	gameint, _ := strconv.Atoi(gameID)

	// Consultar Redis para ver si los detalles del juego están en caché.
	cachedDetails, err := gc.GetCachedGameDetails(gameint)
	if err != nil {
		log.Printf("Error al consultar detalles en caché: %v", err)
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if cachedDetails != nil {
		// Si los detalles están en caché, responder con los detalles en caché.
		ctx.JSON(200, cachedDetails)
		return
	}

	// Intentar obtener los detalles de la base de datos.
	dbDetails, dbErr := gc.dbClient.GetGameDetails(gameint)
	if dbErr != nil && !errors.Is(dbErr, sql.ErrNoRows) {
		ctx.JSON(400, gin.H{
			"Error al obtener detalles del juego desde la base de datos": dbErr.Error(),
		})
		log.Printf("Error al obtener detalles de la base de datos: %v", dbErr)
		return

	}
	if dbDetails != nil {
		// Si los detalles están en la base de datos, responder con los detalles de la base de datos.
		ctx.JSON(200, dbDetails)
		// Guardar los detalles obtenidos de la base de datos en caché.
		err := gc.SaveToCache(gameID, dbDetails)
		if err != nil {
			log.Printf("Error al guardar detalles del juego en caché: %v", err)
			ctx.JSON(400, gin.H{"error:": err.Error()})
		}
		return
	}

	// Si no se encontraron detalles en la base de datos, obtenerlos y procesarlos.
	responseData, err := gc.fetchAndProcessGameDetails(gameint)
	if err != nil {
		log.Printf("Error al obtener y procesar detalles del juego: %v", err)
		// Manejar errores si ocurren durante la obtención y procesamiento.
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = gc.dbClient.SaveGameDetails(responseData)
	if err != nil {
		log.Printf("Error al guardar detalles del juego en la base de datos: %v", err)
		ctx.JSON(404, gin.H{"error:": err.Error()})
	}

	// Responder con los detalles obtenidos y procesados.
	ctx.JSON(200, responseData)
}

func (gc *GameControllers) fetchAndProcessGameDetails(gameint int) ([]steamapi.AppDetails, error) {
	// Obtener los detalles del juego de la API de Steam.
	apiDetails, err := gc.steamService.GetAppDetails(gameint)
	if err != nil {
		log.Printf("Error al obtener detalles de la API de Steam: %v", err)
		return nil, err
	}

	// Obtener los detalles de los juegos de la base de datos.
	games, err := gc.dbClient.GetAllByAppID(gameint)
	if err != nil {
		log.Printf("Error al obtener detalles de la base de datos: %v", err)
		return nil, err
	}

	// Procesar los detalles de la API y los detalles de los juegos de la base de datos.
	apiDetailsSlice := [][]byte{apiDetails}
	responseData, err := gc.gameProcessor.ProcessResponse(apiDetailsSlice, games)
	if err != nil {
		log.Printf("Error al procesar detalles del juego: %v", err)
		return nil, err
	}

	return responseData, nil
}

func (gc *GameControllers) GetCachedGameDetails(gameID int) (*entity.GameDetails, error) {
	cachedDetails, err := gc.redisClient.Get(strconv.Itoa(gameID))
	if err != nil && err != redis.Nil {
		log.Printf("Error al consultar detalles en caché: %v", err)
		return cachedDetails, err
	}

	return cachedDetails, nil
}

func encodeToJSON(data interface{}) ([]byte, error) {
	encodedData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error al codificar datos a JSON: %v", err)
		return encodedData, err
	}
	return encodedData, nil
}

func (gc *GameControllers) SaveToCache(gameID string, data interface{}) error {
	// Codifica los detalles en formato JSON.
	encodedData, err := encodeToJSON(data)
	if err != nil {
		log.Printf("Error al codificar detalles del juego a JSON: %v", err)
		return err
	}

	// Guarda los detalles en caché utilizando Redis.
	err = gc.redisClient.Set(gameID, string(encodedData))
	if err != nil {
		log.Printf("Error al guardar detalles del juego en caché: %v", err)
		return err
	}

	return nil
}
