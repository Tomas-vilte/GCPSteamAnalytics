package config

import (
	"log"
	"os"
)

type Config struct {
	DBUser                 string
	DBPass                 string
	DBName                 string
	InstanceConnectionName string
}

func LoadEnvVariables() *Config {
	config := &Config{
		DBUser:                 setEnvVariable("DB_USER"),
		DBPass:                 setEnvVariable("DB_PASS"),
		DBName:                 setEnvVariable("DB_NAME"),
		InstanceConnectionName: setEnvVariable("INSTANCE_CONNECTION_NAME"),
	}
	return config
}

func setEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Error: %s variable de entorno no establecida", key)
	}
	return value
}
