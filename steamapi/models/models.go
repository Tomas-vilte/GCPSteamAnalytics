package models

import "context"

// AppDetails representa los detalles de una aplicación en la tienda Steam.
type AppDetails struct {
	SteamAppid            int64    `json:"steam_appid"`
	Type                  string   `json:"type"`
	Name                  string   `json:"name"`
	Description           string   `json:"short_description"`
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
	PriceOverview struct {
		Currency        string `json:"currency"`
		DiscountPercent int64  `json:"discount_percent"`
		Initial         int64  `json:"initial"`
		FinalFormatted  string `json:"final_formatted"`
	} `json:"price_overview"`
}

// AppDetailsResponse es la estructura de respuesta para los detalles de la aplicación.
type AppDetailsResponse struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

// SteamData es una interfaz que define los métodos para interactuar con los datos de Steam.
type SteamData interface {
	// ProcessSteamData obtiene los detalles de las aplicaciones Steam en paralelo.
	// Utiliza la información de contexto 'ctx' para controlar la ejecución y
	// procesa hasta 'limit' aplicaciones a partir de los IDs de aplicación 'appIDs'.
	ProcessSteamData(ctx context.Context, appIDs []int, limit int) ([]AppDetails, error)

	// ProcessAppID obtiene los detalles de una aplicación Steam específica.
	// Devuelve un puntero a AppDetails que contiene los detalles de la aplicación
	// correspondiente al ID proporcionado 'id'. Si hay un error, se devuelve junto con nil.
	ProcessAppID(id int64) (*AppDetails, error)

	// GetAllAppIDs obtiene todos los appIDs almacenados en la base de datos MySQL.
	// Devuelve una lista de IDs de aplicaciones y un posible error en caso de fallo.
	GetAllAppIDs(limit int) ([]int, error)

	UpdateAppStatus(id int, isValid bool) error

	// SaveToCSV guarda los detalles de las aplicaciones en un archivo CSV.
	// 'data' es una lista de AppDetails que se guardarán en el archivo especificado por 'filePath'.
	// Devuelve un posible error en caso de fallo al guardar los datos en el archivo.
	SaveToCSV(data []AppDetails, filePath string) error
}
