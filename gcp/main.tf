provider "google" {
  credentials = file("/home/tomi/GCPSteamAnalytics/application_default_credentials.json")
  project     = "gcpsteamanalytics"
  region      = "US"
}

module "cloud_function" {
  source = "./cloud_function"
}

module "api_gateway" {
  source = "./api_gateway"
}
  