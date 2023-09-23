provider "google" {
  project = "gcpsteamanalytics"
  region  = "US"
}

# module "cloud_storage" {
#   source = "./cloud_storage"
# }

# module "my_network" {
#   source = "./vpc"
# }

# module "api_gateway" {
#   source = "./api_gateway"
# }
  
# module "cloud_function" {

#   DB_USER = var.DB_USER
#   DB_PASS = var.DB_PASS
#   DB_NAME = var.DB_NAME
#   INSTANCE_CONNECTION_NAME = var.INSTANCE_CONNECTION_NAME
#   REDISHOST = var.REDISHOST
#   source = "./cloud_function" 
# }

# module "cloud_sql" {
#   source = "./cloud_sql"
# }

# module "redis" { 
#   source = "./memory_storage_redis"

# }