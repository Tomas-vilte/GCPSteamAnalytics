package models

type BestPurchaseOption struct {
	Id_purchase       int64 `json:"id_purchase"`
	Id_storeitem      int64 `json:"id_storeitem"`
	Discount_pct      int64 `json:"discount_pct"`
	Discount_amount   int64 `json:"discount_amount"`
	Discount_end_date int64 `json:"discount_end_date"`
	Price_original    int64 `json:"price_original"`
}
