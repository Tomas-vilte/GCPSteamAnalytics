package entity

import (
	"errors"
	"strings"
)

type LanguageArray []string

func (l *LanguageArray) Scan(src interface{}) error {
	// Convierte el valor de la base de datos ([]uint8) a []byte.
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("no se pudo convertir a []byte")
	}

	// Divide los bytes en cadenas usando una coma como delimitador.
	languages := strings.Split(string(bytes), ",")

	// Asigna el resultado a la variable.
	*l = languages

	return nil
}

type SupportLanguages struct {
	InterfaceLanguages LanguageArray `json:"interface_languages" db:"interface_languages"`
	FullAudioLanguages LanguageArray `json:"full_audio_languages" db:"fullAudio_languages"`
	SubtitlesLanguages LanguageArray `json:"subtitles_languages" db:"subtitles_languages"`
}

type Platforms struct {
	Windows bool `json:"windows" db:"windows"`
	Mac     bool `json:"mac" db:"mac"`
	Linux   bool `json:"linux" db:"linux"`
}

type ReleaseDate struct {
	Date       string `json:"release_date" db:"release_date"`
	ComingSoon bool   `json:"coming_soon" db:"coming_soon"`
}

type Price struct {
	Currency         string `json:"currency" db:"currency"`
	DiscountPercent  int    `json:"discount_percent" db:"discount_percent"`
	InitialFormatted string `json:"initial_formatted" db:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted" db:"final_formatted"`
}

type GameDetails struct {
	ID               int    `json:"id" db:"id"`
	AppID            int    `json:"app_id" db:"app_id"`
	Description      string `json:"description" db:"description"`
	Type             string `json:"type" db:"type"`
	Name             string `json:"name" db:"name"`
	Publishers       string `json:"publishers" db:"publishers"`
	Developers       string `json:"developers" db:"developers"`
	IsFree           bool   `json:"is_free" db:"is_free"`
	SupportLanguages `json:"support_languages"`
	Platforms        `json:"platforms"`
	ReleaseDate      `json:"release_date"`
	Price            `json:"price"`
}

//type GameDetails struct {
//	ID                 int    `db:"id"`
//	AppID              int    `db:"app_id"`
//	Description        string `db:"description"`
//	Type               string `db:"type"`
//	Name               string `db:"name"`
//	Publishers         string `db:"publishers"`
//	Developers         string `db:"developers"`
//	IsFree             bool   `db:"is_free"`
//	InterfaceLanguages string `db:"interface_languages"`
//	FullAudioLanguages string `db:"fullAudio_languages"`
//	SubtitlesLanguages string `db:"subtitles_languages"`
//	Windows            bool   `db:"windows"`
//	Mac                bool   `db:"mac"`
//	Linux              bool   `db:"linux"`
//	ReleaseDate        string `db:"release_date"`
//	ComingSoon         bool   `db:"coming_soon"`
//	Currency           string `db:"currency"`
//	DiscountPercent    int    `db:"discount_percent"`
//	InitialFormatted   string `db:"initial_formatted"`
//	FinalFormatted     string `db:"final_formatted"`
//}
