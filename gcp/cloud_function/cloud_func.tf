resource "google_cloudfunctions_function" "cloud_function" {
  name        = "get_data" 
  runtime     = "go120"              
  source_archive_bucket = "steam-analytics"
  source_archive_object = "steamAPI.zip" 
  entry_point = "ProcessSteamDataAndSaveToStorage"
  trigger_http = true

 
}