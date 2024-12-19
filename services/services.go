// Handles the API call and processes the IP location data
package services

import (
	"encoding/json"
	"fmt"
	"ip_locator/entities"
	"net/http"
)

const apiURL = "http://ip-api.com/json/"

// GetIPLocation fetches location data for a given IP address
func GetIPLocation(ip string) (*entities.IPLocation, error) {
	// Build the API URL
	url := fmt.Sprintf("%s%s", apiURL, ip)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IP location: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var location entities.IPLocation
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Check the status field in the API response
	if location.Status != "success" {
		return nil, fmt.Errorf("failed to locate IP '%s': status %s", ip, location.Status)
	}

	// Return the location data
	return &location, nil
}
