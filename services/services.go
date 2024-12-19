package services

import (
	"encoding/json"
	"fmt"
	"ip_locator/entities"
	"net/http"
)

const apiURL = "http://ip-api.com/json/"

func GetIPLocation(ip string) (*entities.IPLocation, error) {
	url := fmt.Sprintf("%s%s", apiURL, ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IP location: %w", err)
	}
	defer resp.Body.Close()

	var location entities.IPLocation
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &location, nil
}