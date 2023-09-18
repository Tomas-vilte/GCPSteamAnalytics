cd ..
zip -r my_functions/process_games/process_games.zip process_games.go go.mod go.sum db controller api steamapi/persistence steamapi/service steamapi/model handlers cache utils config configGCP.env

gsutil cp my_functions/process_games/process_games.zip gs://steam-analytics

gcloud functions deploy process-games --region=us-central1 --source=gs://steam-analytics/process_games.zip