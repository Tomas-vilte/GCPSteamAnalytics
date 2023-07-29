resource "google_storage_bucket" "bucket" {
  name = "steam-analytics"
  location = "US"
  storage_class = "STANDARD"
  project = "gcpsteamanalytics"
  

  uniform_bucket_level_access = true
}
