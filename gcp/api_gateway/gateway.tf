resource "google_api_gateway_api" "api_gateway" {
  provider = google-beta
  api_id   = "api_gateway"
  display_name = "api_gateway"
  openapi_config {
    spec = file("${path.module}/openapi-functions.yaml")
  }
}