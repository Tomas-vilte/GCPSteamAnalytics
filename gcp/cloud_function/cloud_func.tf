resource "google_cloudfunctions2_function" "function1" {
  name        = "process-games"
  location    = "us-central1"
  description = "steam-analytics"
  build_config {
    runtime     = "go120"
    entry_point = "MyCloudFunction" 
    source {
      storage_source {  
        bucket = "steam-analytics"
        object = "process_games.zip"
      }
    }
  }
  service_config {
    max_instance_count = 1
    available_memory   = "1024M"
    timeout_seconds    = 60
    ingress_settings = "ALLOW_ALL"
    vpc_connector = "my-vpc"
    vpc_connector_egress_settings = "PRIVATE_RANGES_ONLY"
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
    max_instance_count = 1
    available_memory = "1024M"
    timeout_seconds = 60
    ingress_settings = "ALLOW_ALL"
    vpc_connector = "my-vpc"
    vpc_connector_egress_settings = "PRIVATE_RANGES_ONLY"
  }
}