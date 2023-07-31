package models

type SupportLanguage struct {
	Id_language  int64 `json:"id_language"`
	Id_storeitem int64 `json:"id_storeitem"`
	Full_audio   bool  `json:"full_audio"`
	Subtitles    bool  `json:"subtitles"`
	Supported    bool  `json:"supported"`
}
