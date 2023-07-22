package main

import (
	"fmt"
	"steamAPI/api/config"
	"steamAPI/api/utilities"
)

func main() {
	fmt.Println(config.GetCrendentials())
	file := "./data/hola2.txt"
	fmt.Println(utilities.UploadFileToGCS(file, "steam-analytics", "hola12.txt"))

}
