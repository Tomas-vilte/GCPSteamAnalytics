variable "DB_USER" {
  description = "Nombre de usuario de la base de datos"
  type        = string
}

variable "DB_PASS" {
  description = "Contraseña de la base de datos"
  type        = string
}

variable "DB_NAME" {
  description = "Nombre de la base de datos"
  type        = string
}

variable "INSTANCE_CONNECTION_NAME" {
  description = "Nombre de conexión de la instancia de Cloud SQL"
  type        = string
}

variable "REDISHOTS" {
  description = "Dirrecion host de MemoryStore de Redis"
  type = string
}