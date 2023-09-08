package controller

import (
	steamapi "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/model"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func (gc *GameControllers) GetGames(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageNum := parseParam(page)
	pageSizeNum := parseParam(pageSize)

	startIndex := (pageNum - 1) * pageSizeNum

	games, totalItems, err := gc.dbClient.GetGamesByPage(startIndex, pageSizeNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metadata := generateMetadata(totalItems, pageNum, pageSizeNum)
	response := steamapi.PaginatedResponse{
		Metadata: metadata,
		Games:    games,
	}

	c.JSON(http.StatusOK, response)
}

func parseParam(param string) int {
	num, err := strconv.Atoi(param)
	if err != nil || num <= 0 {
		return 1
	}
	return num
}

func generateMetadata(total, page, size int) map[string]interface{} {
	totalPages := int(math.Ceil(float64(total) / float64(size)))

	metadata := map[string]interface{}{
		"total":       total,
		"page":        page,
		"size":        size,
		"total_pages": totalPages,
	}

	return metadata
}
