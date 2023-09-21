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
	languages := strings.Split(string(bytes), ", ")

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

type FullGame struct {
	FullGameAppID int    `json:"appid" db:"fullgame_app_id"`
	FullGameName  string `json:"fullgame_name" db:"fullgame_name"`
}

type Genre struct {
	GenreID   string `json:"genre_id" db:"genre_id"`
	TypeGenre string `json:"type_genre" db:"type_genre"`
}

type Price struct {
	Currency              string  `json:"currency" db:"currency"`
	InitialPrice          float64 `json:"initial_price" db:"initial_price"`
	FinalPrice            float64 `json:"final_price" db:"final_price"`
	DiscountPercent       int     `json:"discount_percent" db:"discount_percent"`
	FormattedInitialPrice string  `json:"formatted_initial_price" db:"formatted_initial_price"`
	FormattedFinalPrice   string  `json:"formatted_final_price" db:"formatted_final_price"`
}

type GameDetails struct {
	ID               int    `json:"id" db:"id"`
	AppID            int    `json:"app_id" db:"app_id"`
	Name             string `json:"name" db:"name"`
	Description      string `json:"description" db:"description"`
	FullGame         `json:"fullgame"`
	Type             string        `json:"type" db:"type"`
	Publishers       LanguageArray `json:"publishers" db:"publishers"`
	Developers       LanguageArray `json:"developers" db:"developers"`
	IsFree           bool          `json:"is_free" db:"is_free"`
	SupportLanguages `json:"support_languages"`
	Platforms        `json:"platforms"`
	Genre            []Genre `json:"genre"`
	ReleaseDate      `json:"release_date"`
	Price            `json:"price"`
	GenreID          string `json:"-" db:"genre_id"`
	TypeGenre        string `json:"-" db:"type_genre"`
}
