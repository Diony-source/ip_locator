// Defines the structure for IP location data
package entities

// IPLocation represents the response structure from the API
type IPLocation struct {
	IP          string  `json:"query"`
	Country     string  `json:"country"`
	Region      string  `json:"regionName"`
	City        string  `json:"city"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	ISP         string  `json:"isp"`
	Organization string `json:"org"`
	Status      string  `json:"status"`
}
