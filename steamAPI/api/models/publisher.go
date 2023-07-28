package models

type Publisher struct {
	Id_publisher int64  `json:"id_publisher"`
	Id_storeitem int64  `json:"id_storeitem"`
	Name         string `json:"name"`
}
