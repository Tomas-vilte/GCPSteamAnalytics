package main

import (
	"fmt"
	"steamAPI/api/handlers"
)

func main() {
	//fmt.Println(config.GetCrendentials())
	//file := "/home/tomi/GCPSteamAnalytics/SteamAPI/api/data/hola2.txt"
	//fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))

	res, _ := handlers.GetData()
	for _, v := range res {
		fmt.Println(v)
	}
	fmt.Println(handlers.GetData())

}
