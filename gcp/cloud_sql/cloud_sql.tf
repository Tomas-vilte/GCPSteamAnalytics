resource "google_compute_instance" "my_instance" {
  name         = "my-instance"
  machine_type = "e2-micro"
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
  database_version = "MYSQL_5_7"
  region           = "us-central1"
  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_user" "my_db_user" {
  name     = "my-db-user"
  instance = google_sql_database_instance.my_db_instance.name
  password = "root"
}

resource "google_sql_database" "my_db" {
  name     = "my-database"
  instance = google_sql_database_instance.my_db_instance.name
}
