package models

type Plataform struct {
	Id_game      int64 `json:"id_game"`
	Id_storeitem int64 `json:"id_storeitem"`
	Windows      bool  `json:"windows"`
	Mac          bool  `json:"mac"`
	Linux        bool  `json:"liunx"`
}
