package controller

import (
	"fmt"
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func (gc *GameControllers) GetGames(c *gin.Context) {
	games, err := gc.dbClient.GetAllGames()
	if err != nil {
		c.JSON(400, gin.H{
			"Error al obtener los juegos:": err.Error(),
		})
		return
	}
	gc.PaginateGames(c, games, 5)
}

func (gc *GameControllers) PaginateGames(c *gin.Context, games []entity.GameDetails, pageSize int) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", strconv.Itoa(pageSize))

	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	cacheKey := generateCacheKey(page, size)
	cachedResponse, err := gc.redisClient.GetGames(cacheKey)

	if err == nil {
		c.JSON(200, cachedResponse)
		return
	}

	paginatedGames := getPaginatedGames(games, pageInt, sizeInt)

	metadata := generateMetadata(len(games), pageInt, sizeInt)
	response := steamapi.PaginatedResponse{
		Metadata: metadata,
		Games:    paginatedGames,
	}

	err = gc.SaveToCache(cacheKey, response)
	if err != nil {
		c.JSON(404, gin.H{
			"Error al cachear los juegos:": err,
		})
	}

	c.JSON(http.StatusOK, response)
}

func generateCacheKey(page, size string) string {
	return fmt.Sprintf("games_page_%s_size_%s", page, size)
}

func getPaginatedGames(games []entity.GameDetails, page, size int) []entity.GameDetails {
	start := (page - 1) * size
	end := start + size

	if end > len(games) {
		end = len(games)
	}

	return games[start:end]
}

func generateMetadata(total, page, size int) map[string]interface{} {
	return map[string]interface{}{
		"total":       total,
		"page":        page,
		"size":        size,
		"total_pages": int(math.Ceil(float64(total) / float64(size))),
	}
}
