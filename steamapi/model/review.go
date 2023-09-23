package model

type ReviewSummary struct {
	NumReviews int `json:"num_reviews"`
}

type ReviewResponse struct {
	ReviewSummary ReviewSummary `json:"query_summary"`
	Success       int           `json:"success"`
	Reviews       []Review      `json:"reviews"`
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
	RecommendationID         string       `json:"recommendationid"`
	Author                   ReviewAuthor `json:"author"`
	Language                 string       `json:"language"`
	ReviewText               string       `json:"review"`
	TimestampCreated         int          `json:"timestamp_created"`
	TimestampUpdated         int          `json:"timestamp_updated"`
	VotedUp                  bool         `json:"voted_up"`
	VotesUp                  int          `json:"votes_up"`
	VotesFunny               int          `json:"votes_funny"`
	CommentCount             int          `json:"comment_count"`
	SteamPurchase            bool         `json:"steam_purchase"`
	ReceivedForFree          bool         `json:"received_for_free"`
	WrittenDuringEarlyAccess bool         `json:"written_during_early_access"`
}

type ReviewsDB struct {
	AppID                    int    `json:"app_id" db:"app_id"`
	ReviewType               string `json:"review_type" db:"review_type"`
	RecommendationID         int    `json:"recommendation_id" db:"recommendation_id"`
	SteamID                  string `json:"steam_id" db:"steam_id"`
	NumGamesOwned            int    `json:"num_games_owned" db:"num_games_owned"`
	NumReviews               int    `json:"num_reviews" db:"num_reviews"`
	PlaytimeForever          int    `json:"playtime_forever" db:"playtime_forever"`
	PlaytimeLastTwoWeeks     int    `json:"playtime_last_two_weeks" db:"playtime_last_two_weeks"`
	PlaytimeAtReview         int    `json:"playtime_at_review" db:"playtime_at_review"`
	LastPlayed               int    `json:"last_played" db:"last_played"`
	Language                 string `json:"language" db:"language"`
	ReviewText               string `json:"review_text" db:"review_text"`
	TimestampCreated         int    `json:"timestamp_created" db:"timestamp_created"`
	TimestampUpdated         int    `json:"timestamp_updated" db:"timestamp_updated"`
	VotedUp                  bool   `json:"voted_up" db:"voted_up"`
	VotesUp                  int    `json:"votes_up" db:"votes_up"`
	VotesFunny               int    `json:"votes_funny" db:"votes_funny"`
	CommentCount             int    `json:"comment_count" db:"comment_count"`
	SteamPurchase            bool   `json:"steam_purchase" db:"steam_purchase"`
	ReceivedForFree          bool   `json:"received_for_free" db:"received_for_free"`
	WrittenDuringEarlyAccess bool   `json:"written_during_early_access" db:"written_during_early_access"`
}
