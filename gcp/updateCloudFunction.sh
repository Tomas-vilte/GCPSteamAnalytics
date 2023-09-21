cd ..

zip -r my_functions/process_games/process_games.zip process_games.go go.mod go.sum controller/process_controller.go api/routes_process_game.go steamapi/persistence/entity steamapi/persistence/client_db.go steamapi/persistence/storage.go steamapi/service/process_data.go steamapi/service/steam_client.go steamapi/model/app_detail.go utils/steam_utils.go config
zip -r my_functions/get_games/get_games.zip games_from_db.go go.mod go.sum controller/game_appid_controller.go controller/game_controller.go api/routes_get_games.go steamapi/persistence/client_db.go steamapi/persistence/storage.go steamapi/service/process_data.go steamapi/service/steam_client.go steamapi/model/app_detail.go steamapi/persistence/entity cache/redis_cache.go utils/steam_utils.go config

gsutil cp my_functions/process_games/process_games.zip gs://steam-analytics
gsutil cp my_functions/get_games/get_games.zip gs://steam-analytics

gcloud functions deploy process-games --region=us-central1 --source=gs://steam-analytics/process_games.zip
gcloud functions deploy get-games --region=us-central1 --source=gs://steam-analytics/get_games.zip