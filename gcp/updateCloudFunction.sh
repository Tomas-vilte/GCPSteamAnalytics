cd ..
zip -r cloud_function/steam.zip go.mod main.go go.sum db controller api steamapi

gsutil cp steam.zip gs://steam-analytics

gcloud functions deploy process-steam-analytics --region=us-central1 --source=gs://steam-analytics/steam.zip