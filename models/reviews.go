package models

type Reviews struct {
	Id_review          int64  `json:"id_review"`
	Id_storeitem       int64  `json:"id_storeitem"`
	Percent_positive   int64  `json:"percent_positive"`
	Review_count       int64  `json:"review_count"`
	Review_score       int64  `json:"review_score"`
	Review_score_label string `json:"review_score_label"`
}
