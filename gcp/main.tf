provider "google" {
  project = "gcpsteamanalytics"
  region  =  "US"
}

module "cloud_storage" {
  source = "./cloud_storage"
}

module "cloud_function" {
  source = "./cloud_function"
}

module "api_gateway" {
  source = "./api_gateway"
}
  
module "cloud_sql" {
  source  = "./cloud_sql"
}