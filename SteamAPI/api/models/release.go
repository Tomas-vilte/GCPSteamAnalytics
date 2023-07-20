package models

type Release struct {
	Id_release         int64 `json:"id_release"`
	Id_storeitem       int64 `json:"id_storeitem"`
	Steam_release_date int64 `json:"steam_release_date"`
}
