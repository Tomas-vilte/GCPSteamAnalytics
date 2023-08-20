package models

type ReviewData interface {
	GetReviews(appID int) (*ReviewResponse, error)
}

type ReviewSummary struct {
	NumReviews      int    `json:"num_reviews"`
	ReviewScore     int    `json:"review_score"`
	ReviewScoreDesc string `json:"review_score_desc"`
	TotalPositive   int    `json:"total_positive"`
	TotalNegative   int    `json:"total_negative"`
	TotalReviews    int    `json:"total_reviews"`
}

type ReviewResponse struct {
	Success int      `json:"success"`
	Reviews []Review `json:"reviews"`
}

type ReviewAuthor struct {
	SteamID              string `json:"steamid"`
	NumGamesOwned        int    `json:"num_games_owned"`
	NumReviews           int    `json:"num_reviews"`
	PlaytimeForever      int    `json:"playtime_forever"`
	PlaytimeLastTwoWeeks int    `json:"playtime_last_two_weeks"`
	PlaytimeAtReview     int    `json:"playtime_at_review"`
	LastPlayed           int    `json:"last_played"`
}

type Review struct {
	ReviewSummary            ReviewSummary `json:"review_summary"`
	RecommendationID         string        `json:"recommendationid"`
	Author                   ReviewAuthor  `json:"author"`
	Language                 string        `json:"language"`
	ReviewText               string        `json:"review"`
	TimestampCreated         int           `json:"timestamp_created"`
	TimestampUpdated         int           `json:"timestamp_updated"`
	VotedUp                  bool          `json:"voted_up"`
	VotesUp                  int           `json:"votes_up"`
	VotesFunny               int           `json:"votes_funny"`
	CommentCount             int           `json:"comment_count"`
	SteamPurchase            bool          `json:"steam_purchase"`
	ReceivedForFree          bool          `json:"received_for_free"`
	WrittenDuringEarlyAccess bool          `json:"written_during_early_access"`
}
