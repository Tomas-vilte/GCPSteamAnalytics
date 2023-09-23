locals {
  api_config_id_prefix     = "api"
  api_gateway_container_id = "api-gw"
  gateway_id               = "gw"
  project_id               = "gcpsteamanalytics"
}

resource "google_api_gateway_api" "api_gw" {
  provider     = google-beta
  api_id       = local.api_gateway_container_id
  project = local.project_id
  display_name = "The API Gateway"

}

resource "google_api_gateway_api_config" "api_cfg" {
  project = local.project_id
  provider      = google-beta
  api           = google_api_gateway_api.api_gw.api_id
  api_config_id_prefix = local.api_config_id_prefix
  display_name  = "The Config"

  openapi_documents {
    document {
      path     = "api_gateway/openapi-functions.yaml"
      contents = filebase64("api_gateway/openapi-functions.yaml")
    }
  }
}

resource "google_api_gateway_gateway" "gw" {
  project    = local.project_id
  provider   = google-beta
  region     = "us-central1"

  api_config   = google_api_gateway_api_config.api_cfg.id

  gateway_id   = local.gateway_id
  display_name = "SteamAPIGateway"

  depends_on = [google_api_gateway_api_config.api_cfg]
}