cd ..
zip -r steam.zip go.mod main.go go.sum utilities handlers functionGCP


gsutil cp steam.zip gs://steam-analytics

gcloud functions deploy process-steam-analytics --region=us-central1 --source=gs://steam-analytics/steam.zip --entry-point=ProcessSteamDataAndSaveToStorage