package steamapi

type AppDetails struct {
	SteamAppid  int64    `json:"steam_appid"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"short_description"`
	Developers  []string `json:"developers"`
	Publishers  []string `json:"publishers"`
	IsFree      bool     `json:"is_free"`
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	PriceOverview struct {
		Currency        string `json:"currency"`
		DiscountPercent int64  `json:"discount_percent"`
		Initial         int64  `json:"initial"`
		FinalFormatted  string `json:"final_formatted"`
	} `json:"price_overview"`
}

type AppDetailsResponse struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

type SteamData interface {
	GetSteamData(appIDs []int, limit int) ([]AppDetails, error)
	GetAllAppIDs(startID int) ([]int, error)
	LoadLastProcessedAppid() (int, error)
	SaveLastProcessedAppid(lastProcessedAppid int) error
}
