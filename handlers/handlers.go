// Manages user interaction for the IP locator app
package handlers

import (
	"fmt"
	"ip_locator/services"
)

func Start() {
	var ip string

	// Get IP address from the user
	fmt.Print("Enter the IP address to locate: ")
	fmt.Scanln(&ip)

	// Call the service to get IP location
	location, err := services.GetIPLocation(ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the location information
	fmt.Println("\nIP Location Information:")
	fmt.Printf("IP: %s\n", location.IP)
	fmt.Printf("Country: %s\n", location.Country)
	fmt.Printf("Region: %s\n", location.Region)
	fmt.Printf("City: %s\n", location.City)
	fmt.Printf("Latitude: %f\n", location.Latitude)
	fmt.Printf("Longitude: %f\n", location.Longitude)
	fmt.Printf("ISP: %s\n", location.ISP)
	fmt.Printf("Organization: %s\n", location.Organization)
}
