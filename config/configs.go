package config

import "os"

func GetCrendentials() string {
	return os.Getenv("")
}
