package steamapi

import (
	"encoding/json"
)

type PriceOverview struct {
	Currency         string `json:"currency"`
	DiscountPercent  int    `json:"discount_percent"`
	InitialFormatted string `json:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted"`
}

type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

type Release struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type StoreItem struct {
	SteamAppid       int64  `json:"steam_appid"`
	NameGame         string `json:"name"`
	ShortDescription string `json:"short_description"`
	IsFree           bool   `json:"is_free"`
}

type GameDetails struct {
	StoreItem      StoreItem
	Developers     []string      `json:"developers"`
	Publishers     []string      `json:"publishers"`
	PriceOverviews PriceOverview `json:"price_overview"`
	Platform       Platforms     `json:"platforms"`
	ReleaseDate    Release       `json:"release_date"`
	Type           string        `json:"type"`
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
