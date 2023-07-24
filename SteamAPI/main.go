package main

import (
	"fmt"
	"steamAPI/api/config"
	"steamAPI/api/utilities"
)

func main() {
	fmt.Println(config.GetCrendentials())
	file := "/home/tomi/GCPSteamAnalytics/SteamAPI/api/data/hola2.txt"
	fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))

}
