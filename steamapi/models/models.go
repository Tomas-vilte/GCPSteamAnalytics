package models

// AppDetails representa los detalles de una aplicación en la tienda Steam.
type AppDetails struct {
	SteamAppid  int64    `json:"steam_appid"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"short_description"`
	Developers  []string `json:"developers"`
	Publishers  []string `json:"publishers"`
	IsFree      bool     `json:"is_free"`
	// ReleaseDate contiene información sobre la fecha de lanzamiento.
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	// Platforms contiene información sobre las plataformas compatibles.
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	// PriceOverview contiene detalles sobre el precio de la aplicación.
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
	ProcessSteamData(appIDs []int, limit int) ([]AppDetails, error)
	// ProcessAppID obtiene los detalles de una aplicación Steam específica.
	ProcessAppID(id int) (*AppDetails, error)
	// GetAllAppIDs obtiene todos los appIDs almacenados en la base de datos MySQL.
	GetAllAppIDs(lastProcessedAppID int) ([]int, error)
	// LoadLastProcessedAppid carga el último appID procesado de la base de datos.
	LoadLastProcessedAppid() (int, error)
	// SaveLastProcessedAppid guarda el último appID procesado en la base de datos.
	SaveLastProcessedAppid(lastProcessedAppid int) error
	// SaveToCSV guarda los detalles de las aplicaciones en un archivo CSV.
	SaveToCSV(data []AppDetails, filePath string) error
}
