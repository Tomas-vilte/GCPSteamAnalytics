package models

type StoreItem struct {
	Id_game           int64  `json:"id_game"`
	Name              string `json:"name"`
	Short_description string `json:"short_description"`
}
