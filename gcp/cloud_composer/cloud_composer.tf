resource "google_composer_environment" "test" {
  name = "steam-airflow"
  region = "us-central1"
  config {
    
    node_config {
      network = "projects/gcpsteamanalytics/global/networks/my-vpc"
      subnetwork = "projects/gcpsteamanalytics/global/networks/my-vpc"
    }
    software_config {
      image_version = "composer-2.4.3-airflow-2.5.3"
    }
    workloads_config {
      scheduler {
        cpu = 0.5
        memory_gb = 1.875
        storage_gb = 1
        count = 1
      }
    web_server {
      cpu = 0.5
      memory_gb = 1.875
      storage_gb = 1
      
    }
    worker {
      cpu = 0.5
      memory_gb = 1.875
      storage_gb = 1
      max_count = 3
      min_count = 1
    }
    }
    environment_size = "ENVIRONMENT_SIZE_SMALL"
  }  
}