resource "google_redis_instance" "cache" {
    name = "my-redis-instance"
    region = "us-central1"
    tier = "BASIC"
    memory_size_gb = 5
}