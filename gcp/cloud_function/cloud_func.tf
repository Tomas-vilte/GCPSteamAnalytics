// crea una funcion que corra golang 1.20 en la region de us-central1
// el archivo .zip que contiene el codigo de la funcion se encuentra en bucket steam-analytics 
// el nombre del archivo es steam-analytics.zip
// la funcion se llama steam-analytics
// su trigger es un metodo post
// el runtime es go 1.20
// tiene 1gb y 2 nucleos
// tipo de evento google.cloud.apigateway.gateway.v1.created
// las cloud functions tienen que ser de generacion 2

resource "google_cloudfunctions2_function" "function" {
  name        = "process-steam-analytics"
  location    = "us-central1"
  description = "steam-analytics"

  build_config {
    runtime     = "go120"
    entry_point = "ProcessSteamDataAndSaveToStorage" # Set the entry point
    source {
      storage_source {  
        bucket = "steam-analytics"
        object = "steam.zip"
      }
    }
  }

  service_config {
    max_instance_count = 1
    available_memory   = "1024M"
    timeout_seconds    = 60
  }
}