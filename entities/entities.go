// Defines the structure for IP location data
package entities

// IPLocation represents the response structure from the API
type IPLocation struct {
	IP          string `json:"ip"`
	Country     string `json:"country_name"`
	Region      string `json:"region_name"`
	City        string `json:"city"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	ISP         string `json:"isp"`
	Organization string `json:"org"`
}
