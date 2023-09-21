resource "google_cloudfunctions2_function" "function1" {
  name        = "process-games"
  location    = "us-central1"
  description = "steam-analytics"
  build_config {
    runtime     = "go120"
    entry_point = "ProcessGames" 
    source {
      storage_source {  
        bucket = "steam-analytics"
        object = "process_games.zip"
      }
    }
  }
  service_config {
     environment_variables = {
      DB_USER = var.DB_USER
      DB_PASS = var.DB_PASS
      DB_NAME = var.DB_NAME
      INSTANCE_CONNECTION_NAME = var.INSTANCE_CONNECTION_NAME
      REDISHOST = var.REDISHOTS
    }
    max_instance_count = 1
    available_memory   = "1024M"
    timeout_seconds    = 60
    ingress_settings = "ALLOW_ALL"
    vpc_connector = "vpc"
    vpc_connector_egress_settings = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
  }
}

resource "google_cloudfunctions2_function" "function2" {
  name = "get-games"
  location = "us-central1"
  description = "Get games from Cloud SQL"
  build_config {
    runtime = "go120"
    entry_point = "GetGames"
    source {
      storage_source {
        bucket = "steam-analytics"
        object = "get_games.zip"
      }
    }
  }
  service_config {
   environment_variables = {
      DB_USER = var.DB_USER
      DB_PASS = var.DB_PASS
      DB_NAME = var.DB_NAME
      INSTANCE_CONNECTION_NAME = var.INSTANCE_CONNECTION_NAME
      REDISHOST = var.REDISHOTS
    }
    max_instance_count = 1
    available_memory = "1024M"
    timeout_seconds = 60
    ingress_settings = "ALLOW_ALL"
    vpc_connector = "vpc"
    vpc_connector_egress_settings = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
  }
}