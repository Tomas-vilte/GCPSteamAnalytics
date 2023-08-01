// crea una funcion que corra golang 1.20 en la region de us-central1
// el archivo .zip que contiene el codigo de la funcion se encuentra en bucket steam-analytics 
// el nombre del archivo es steam-analytics.zip
// la funcion se llama steam-analytics
// su trigger es un metodo post
// el runtime es go 1.20
// tiene 1gb y 2 nucleos
// tipo de evento google.cloud.apigateway.gateway.v1.created
// las cloud functions tienen que ser de generacion 2


resource "google_cloudfunctions_function" "cloud_function" {
  name        = "process-steam-analytics"
  description = "steam-analytics"
  runtime     = "go120"
  available_memory_mb = 1024
  source_archive_bucket = "steam-analytics"
  source_archive_object = "steam.zip"
  entry_point = "ProcessSteamDataAndSaveToStorage"
  timeout = 60
  trigger_http = true
  region = "us-central1"
  project = "gcpsteamanalytics"
}
