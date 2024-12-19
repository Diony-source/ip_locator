package main

import (
	"fmt"
	"ip_locator/handlers"
)

func main() {
	// Welcome message and start the program
	fmt.Println("Welcome to the IP Locator App!")
	handlers.Start()
}
