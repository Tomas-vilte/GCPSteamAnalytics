resource "google_compute_instance" "my_instance" {
  name         = "my-instance"
  machine_type = "db-n1-standard-4"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }
  

  network_interface {
    network = "default"

  }
  
}

resource "google_sql_database_instance" "my_db_instance" {
  name             = "my-db-instance"
  database_version = "MYSQL_8_0"
  region           = "us-central1"
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled = true
      authorized_networks {
        name = "my-authorized-network"
        value = "181.165.142.76"
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
