resource "google_compute_network" "my_network" {
  name = "my-vpc"
}

resource "google_compute_subnetwork" "my_subnet" {
  name          = "my-subnet"
  network       = google_compute_network.my_network.self_link
  ip_cidr_range = "10.0.0.0/24"  
  region        = "us-central1"
}