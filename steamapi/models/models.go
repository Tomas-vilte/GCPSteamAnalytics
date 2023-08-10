package steamapi

import (
	"encoding/json"
	"github.com/Tomas-vilte/GCPSteamAnalytics/models"
)

type GameDetails struct {
	models.StoreItem
	Developers []string `json:"developers"`
}

type SteamApiResponse struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type SteamData interface {
	ExtractAndSaveLimitedGameDetails(limit int) error
	GetAppIDs(appid int) ([]int, error)
	GameExistsInDatabase(appid int) (bool, error)
	SaveLastProcessedAppid(lastProcessedAppid int) error
	LoadLastProcessedAppid() (int, error)
}
