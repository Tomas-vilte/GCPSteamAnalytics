resource "google_cloudfunctions2_function" "function" {
  name        = "process-steam-analytics"
  location    = "us-central1"
  description = "steam-analytics"

  build_config {
    runtime     = "go120"
    entry_point = "ProcessSteamDataAndSaveToStorage" 
    source {
      storage_source {  
        bucket = "steam-analytics"
        object = "steam.zip"
      }
    }
  }

  service_config {
    max_instance_count = 1
    available_memory   = "1024M"
    timeout_seconds    = 60
  }
}