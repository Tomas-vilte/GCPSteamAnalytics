# Creamos la API en Api gateway
resource "google_api_gateway_api" "api" {
  api_id = "api"
  project = "gcpsteamanalytics"
  location = "us-central1"
}