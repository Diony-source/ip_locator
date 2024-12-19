package handlers

import (
	"fmt"
	"ip_locator/services"
)

func Start() {
	var ip string

	fmt.Print("Enter the IP adress to locate: ")
	fmt.Scanln(&ip)

	location, err := services.GetIPLocation(ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nIP Location Information:")
	fmt.Printf("IP: %s\n", location.IP)
	fmt.Printf("Country: %s\n", location.Country)
	fmt.Printf("Region: %s\n", location.Region)
	fmt.Printf("City: %s\n", location.City)
	fmt.Printf("Latitude: %s\n", location.Latitude)
	fmt.Printf("Longitude: %s\n", location.Longitude)
	fmt.Printf("ISP: %s\n", location.ISP)
	fmt.Printf("Organization: %s\n", location.Organization)
}