package entity

type GameDetails struct {
	AppID            int    `json:"app_id"`
	Description      string `json:"description"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	Publishers       string `json:"publishers"`
	Developers       string `json:"developers"`
	IsFree           bool   `json:"is_free"`
	SupportLanguages struct {
		InterfaceLanguages []string `json:"interface_languages"`
		FullAudioLanguages []string `json:"full_audio_languages"`
		SubtitlesLanguages []string `json:"subtitles_languages"`
	} `json:"support_languages"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	ReleaseDate struct {
		Date       string `json:"date"`
		ComingSoon bool   `json:"coming_soon"`
	} `json:"release_date"`
	Price struct {
		Currency         string `json:"currency"`
		DiscountPercent  int    `json:"discount_percent"`
		InitialFormatted string `json:"initial_formatted"`
		FinalFormatted   string `json:"final_formatted"`
	} `json:"price"`
}
