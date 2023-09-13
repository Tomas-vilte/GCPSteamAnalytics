package model

import "github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"

type Genre struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// AppDetails representa los detalles de una aplicación en la tienda Steam.
type AppDetails struct {
	SteamAppid  int64  `json:"steam_appid"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"short_description"`
	Fullgame    struct {
		AppID string `json:"appid"`
		Name  string `json:"name"`
	} `json:"fullgame"`
	Developers            []string `json:"developers"`
	Publishers            []string `json:"publishers"`
	IsFree                bool     `json:"is_free"`
	SupportedLanguagesRaw string   `json:"supported_languages"`
	SupportedLanguages    map[string][]string
	ReleaseDate           struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	Genres        []Genre `json:"genres"`
	PriceOverview struct {
		Currency         string  `json:"currency"`
		Initial          float64 `json:"initial"`
		Final            float64 `json:"final"`
		DiscountPercent  int64   `json:"discount_percent"`
		InitialFormatted string  `json:"initial_formatted"`
		FinalFormatted   string  `json:"final_formatted"`
	} `json:"price_overview"`
}

// AppDetailsResponse es la estructura de respuesta para los detalles de la aplicación.
type AppDetailsResponse struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

type PaginatedResponse struct {
	Metadata map[string]interface{} `json:"metadata"`
	Games    []entity.GameDetails   `json:"games"`
}
