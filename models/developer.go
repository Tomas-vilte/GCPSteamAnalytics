package models

type Developer struct {
	Id_developer int64  `json:"id_developer"`
	Id_storeitem int64  `json:"id_storeitem"`
	Name         string `json:"name"`
}
