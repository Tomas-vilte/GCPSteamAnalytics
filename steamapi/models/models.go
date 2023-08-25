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
	ProcessAppIDs(appIDs []int, batchSize int) ([]AppDetails, error)

	// GetAllAppIDs obtiene todos los appIDs almacenados en la base de datos MySQL.
	// Devuelve una lista de IDs de aplicaciones y un posible error en caso de fallo.
	GetAllAppIDs(lastProcessedAppID int) ([]int, error)

	// LoadLastProcessedAppid carga el último appID procesado de la base de datos.
	// Devuelve el último appID procesado y un posible error en caso de fallo.
	LoadLastProcessedAppid() (int, error)

	// SaveLastProcessedAppid guarda el último appID procesado en la base de datos.
	// 'lastProcessedAppid' es el último appID que se ha procesado y se debe almacenar.
	// Devuelve un posible error en caso de fallo en la operación.
	SaveLastProcessedAppid(lastProcessedAppid int) error

	// SaveToCSV guarda los detalles de las aplicaciones en un archivo CSV.
	// 'data' es una lista de AppDetails que se guardarán en el archivo especificado por 'filePath'.
	// Devuelve un posible error en caso de fallo al guardar los datos en el archivo.
	SaveToCSV(data []AppDetails, filePath string) error

	// AreEmptyAppIDs verifica si los appIDs dados están presentes en la tabla empty_appids.
	// Devuelve un mapa de IDs de aplicaciones vacías junto con un valor booleano que indica si están vacías.
	// 'appIDs' es una lista de IDs de aplicaciones para verificar.
	AreEmptyAppIDs(appIDs []int) (map[int]bool, error)

	// AddToEmptyAppIDsTable agrega un appID a la tabla empty_appids en la base de datos.
	// 'appID' es el ID de la aplicación que se agregará a la tabla.
	// Devuelve un posible error en caso de fallo en la operación.
	AddToEmptyAppIDsTable(appID int) error

	// GetStartIndexToProcess devuelve el índice de inicio para procesar los appIDs.
	// Utiliza el 'lastProcessedAppID' y la lista de 'appIDs' para determinar el punto
	// de inicio para evitar procesar IDs que ya han sido procesados.
	GetStartIndexToProcess(lastProcessedAppID int, appIDs []int) int
}
