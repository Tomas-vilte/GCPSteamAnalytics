# Creamos la API en Api gateway
resource "google_api_gateway_api" "my_api" {
    api_id = "my-api"
    project = "gcpsteamanalytics"
    location = "us-central1"
}

# Configuramos el endpoint de la API
resource "google_api_gateway_config" "my_api_config" {
    appid = google_api_gateway_api.my_api.id
    project = "gcpsteamanalytics"
    location = "us-central1"
    backend {
        google_cloud_function {
            function = "ProcessSteamDataAndSaveToStorage"
            path = "steamAPI/api/functions"
            
        }
    }
}