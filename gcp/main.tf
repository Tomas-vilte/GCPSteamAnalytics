provider "google" {
    credentials = file("/home/tomi/GCPSteamAnalytics/application_default_credentials.json")
    project = "gcpsteamanalytics"
    region = "US"
}