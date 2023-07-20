package main

import (
	"fmt"
	"steamAPI/api/config"
)

func main() {
	fmt.Println("hola mundo")
	fmt.Println(config.GetCrendentials())

}
