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
  }
}