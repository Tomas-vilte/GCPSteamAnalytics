cd ..

zip -r my_functions/process_games/process_games.zip process_games.go go.mod go.sum db controller api steamapi/persistence steamapi/service steamapi/model handlers cache utils config
zip -r my_functions/get_games/get_games.zip games_from_db.go go.mod go.sum db controller api steamapi/persistence steamapi/service steamapi/model handlers cache utils config

gsutil cp my_functions/process_games/process_games.zip gs://steam-analytics
gsutil cp my_functions/get_games/get_games.zip gs://steam-analytics

gcloud functions deploy process-games --region=us-central1 --source=gs://steam-analytics/process_games.zip
gcloud functions deploy get-games --region=us-central1 --source=gs://steam-analytics/get_games.zip