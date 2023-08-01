# crea un archivo llamada steam.zip que comprima los archivos main.go, go.mod, go.sum y sus carpetas handlers, functionGCP, utilities 
# y lo sube a un bucket de GCP 
# estoy en el path /home/tomi/GCPSteamAnalytics/gcp
# y los archivos y carpetas estan en /home/tomi/GCPSteamAnalytics
# los archivos main.go go.sum, go.mod tienen que estar a primer nivel

# primero crear el .zip
zip -r steam.zip ../main.go ../go.mod ../go.sum ../handlers ../functionGCP ../utilities

# luego subirlo al bucket
gsutil cp steam.zip gs://steam-analytics

# verificar si hay actualizaciones en la infraestructura de GCP con terraform
terraform plan

# actualizar la funcion
gcloud functions deploy process-steam-analytics --region=us-central1 --source=gs://steam-analytics/steam.zip