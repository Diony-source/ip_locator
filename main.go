package main

import (
	"fmt"
	"ip_locator/handlers"
)

func main() {
	fmt.Println("Welcome to the IP Locator App!")
	handlers.Start()
}