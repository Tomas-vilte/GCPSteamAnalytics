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
	GetGameDetails(ctx *gin.Context)
}

type gameController struct {
	steamService  service.SteamClient
	redisClient   cache.RedisClient
	gameProcessor service.GameProcessor
	dbClient      persistence.StorageDB
}

func NewGameController(steamService service.SteamClient, redisClient cache.RedisClient, db persistence.StorageDB, gameProcessor service.GameProcessor) GameController {
	return &gameController{
		steamService:  steamService,
		dbClient:      db,
		gameProcessor: gameProcessor,
		redisClient:   redisClient,
	}
}

func (gc *gameController) GetGameDetails(ctx *gin.Context) {
	gameID := ctx.Param("appid")
	gameint, _ := strconv.Atoi(gameID)

	// Consultar Redis para ver si los detalles del juego están en caché.
	cachedDetails, err := gc.getCachedGameDetails(gameint)
	if err != nil {
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
	dbDetails, dbErr := gc.getDBGameDetails(gameint)
	if dbErr != nil {
		if errors.Is(dbErr, sql.ErrNoRows) {
			// Si no se encontraron detalles en la base de datos, devolver un código de estado 404.
			ctx.JSON(404, gin.H{
				"message": "Juego no encontrado en la base de datos",
			})
			return
		}
		// Manejar otros errores de la base de datos.
		ctx.JSON(500, gin.H{
			"error": dbErr.Error(),
		})
		return
	}

	if dbDetails != nil {
		// Si los detalles están en la base de datos, responder con los detalles de la base de datos.
		ctx.JSON(200, dbDetails)
		// Guardar los detalles obtenidos de la base de datos en caché.
		err := gc.saveToCache(gameID, dbDetails)
		if err != nil {
			// Manejar errores si ocurren durante el almacenamiento en caché.
			log.Printf("Error al guardar detalles del juego en caché: %v", err)
		}
		return
	}

	// Si no se encontraron detalles en la base de datos, obtenerlos y procesarlos.
	responseData, err := gc.fetchAndProcessGameDetails(gameint)
	if err != nil {
		// Manejar errores si ocurren durante la obtención y procesamiento.
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = gc.dbClient.SaveGameDetails(responseData)
	if err != nil {
		// Manejar errores si ocurren durante el almacenamiento en la base de datos.
		log.Printf("Error al guardar detalles del juego en la base de datos: %v", err)
		// Puedes devolver un código de estado 500 aquí si lo deseas.
	}

	// Responder con los detalles obtenidos y procesados.
	ctx.JSON(200, responseData)
}

func (gc *gameController) fetchAndProcessGameDetails(gameint int) ([]steamapi.AppDetails, error) {
	// Obtener los detalles del juego de la API de Steam.
	apiDetails, err := gc.fetchAPIDetails(gameint)
	if err != nil {
		return nil, err
	}

	// Obtener los detalles de los juegos de la base de datos.
	games, err := gc.fetchDBDetails(gameint)
	if err != nil {
		return nil, err
	}

	// Procesar los detalles de la API y los detalles de los juegos de la base de datos.
	responseData, err := gc.processDetails(apiDetails, games)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (gc *gameController) getCachedGameDetails(gameID int) (*entity.GameDetails, error) {
	cachedDetails, err := gc.redisClient.Get(strconv.Itoa(gameID))
	if err != nil && err != redis.Nil {
		return cachedDetails, err
	}

	return cachedDetails, nil
}

func (gc *gameController) getDBGameDetails(gameID int) (*entity.GameDetails, error) {
	dbDetails, err := gc.dbClient.GetGameDetails(gameID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return dbDetails, nil
}

func (gc *gameController) fetchAPIDetails(gameint int) ([]byte, error) {
	// Obtener los detalles del juego de la API de Steam.
	apiDetails, err := gc.steamService.GetAppDetails(gameint)
	if err != nil {
		return nil, err
	}

	return apiDetails, nil
}

func (gc *gameController) fetchDBDetails(gameint int) ([]entity.Item, error) {
	// Obtener los detalles de los juegos de la base de datos.
	games, err := gc.dbClient.GetAllByAppID(gameint)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (gc *gameController) processDetails(apiDetails []byte, games []entity.Item) ([]steamapi.AppDetails, error) {
	// Procesar los detalles de la API y los detalles de los juegos de la base de datos.
	apiDetailsSlice := [][]byte{apiDetails}
	responseData, err := gc.gameProcessor.ProcessResponse(apiDetailsSlice, games)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func encodeToJSON(data interface{}) ([]byte, error) {
	encodedData, err := json.Marshal(data)
	if err != nil {
		return encodedData, err
	}
	return encodedData, nil
}

func (gc *gameController) saveToCache(gameID string, data interface{}) error {
	// Codifica los detalles en formato JSON.
	encodedData, err := encodeToJSON(data)
	if err != nil {
		log.Printf("err: %v\n", err)
		return err
	}

	// Guarda los detalles en caché utilizando Redis.
	err = gc.redisClient.Set(gameID, string(encodedData))
	if err != nil {
		return err
	}

	return nil
}
