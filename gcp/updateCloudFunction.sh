cd ..

zip -r my_functions/process_games/process_games.zip process_games.go go.mod go.sum controller/process_controller.go api/limit.go api/routes_process_game.go steamapi/persistence/entity steamapi/persistence/client_db.go steamapi/persistence/storage.go steamapi/model/review.go steamapi/service/process_data.go steamapi/service/steam_client.go steamapi/model/app_detail.go utils/steam_utils.go config
zip -r my_functions/get_games/get_games.zip games_from_db.go go.mod go.sum controller/game_appid_controller.go controller/game_controller.go api/limit.go api/routes_get_games.go steamapi/persistence/client_db.go steamapi/persistence/storage.go steamapi/model/review.go steamapi/service/process_data.go steamapi/service/steam_client.go steamapi/model/app_detail.go steamapi/persistence/entity cache/redis_cache.go utils/steam_utils.go config
zip -r my_functions/get_reviews/get_review.zip games_reviews.go go.mod go.sum controller/reviews_controller.go api/limit.go api/routes_get_reviews.go steamapi/persistence/client_db.go steamapi/persistence/storage.go steamapi/service/review_fetcher.go steamapi/model steamapi/persistence/entity cache/redis_cache.go config utils/steam_utils.go
zip -r my_functions/api_key/api_key.zip create_api_key.go go.mod go.sum route/route_api_key.go controller/generate_key.go 

gsutil cp my_functions/process_games/process_games.zip gs://steam-analytics
gsutil cp my_functions/get_games/get_games.zip gs://steam-analytics
gsutil cp my_functions/get_reviews/get_review.zip gs://steam-analytics
gsutil cp my_functions/api_key/api_key.zip gs://steam-analytics

gcloud functions deploy process-games-raw --region=us-central1 --source=gs://steam-analytics/process_games.zip
gcloud functions deploy get-games-raw --region=us-central1 --source=gs://steam-analytics/get_games.zip
gcloud functions deploy get-reviews-raw --region=us-central1 --source=gs://steam-analytics/get_reviews.zip
gcloud functions deploy create-api-key --region=us-central1 --source=gs://steam-analytics/api_key.zip