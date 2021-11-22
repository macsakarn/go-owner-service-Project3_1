package main

import (
	"accounts/routes"
	"fmt"
)

func main() {
	fmt.Println("Server Start !")
	routes.StartGin()
}
