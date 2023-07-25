package main

import (
	"fmt"
	"steamAPI/api/handlers"
)

func main() {
	//fmt.Println(config.GetCrendentials())
	//file := "/home/tomi/GCPSteamAnalytics/SteamAPI/api/data/hola2.txt"
	//fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))
	dataFetcher := &handlers.RealDataFetcher{}
	res, err := dataFetcher.GetData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}

}
