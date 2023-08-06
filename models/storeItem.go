package models

type StoreItem struct {
	SteamAppid       int64  `json:"steam_appid"`
	NameGame         string `json:"name"`
	ShortDescription string `json:"short_description"`
}
