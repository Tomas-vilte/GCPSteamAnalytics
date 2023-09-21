resource "google_compute_instance" "my_instance_cloud_sql" {
  name         = "my-instance"
  machine_type = "n1-standard-4"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }
  

  network_interface {
    network = "my-vpc"

  }
  deletion_protection = false
}

resource "google_sql_database_instance" "my_db_instance" {
  name             = "my-db-instance"
  database_version = "MYSQL_8_0"
  region           = "us-central1"
  settings {
    tier = "db-n1-standard-4"
    ip_configuration {
      ipv4_enabled = true  # Habilita la IP pública
      private_network = "projects/gcpsteamanalytics/global/networks/my-vpc"  # Reemplazala con tu VPC
      enable_private_path_for_google_cloud_services = true
      authorized_networks {
        name = "my-authorized-network"
        value = "181.165.142.76" # Aca pone tu ip :D
      }
    }
  }
}

resource "google_sql_user" "my_db_user" {
  name     = "root"
  instance = google_sql_database_instance.my_db_instance.name
  password = "root"
}

resource "google_sql_database" "my_db" {
  name     = "steamAnalytics"
  instance = google_sql_database_instance.my_db_instance.name
}
