package entities

type IPLocator struct {
	IP string `json:"ip"`
	Country string `json:"country"`
	Region string `json:"region"`
	City string `json:"city"`
	Letitude string `json:"letitude"`
	Longitude string `json:"longitude"`
	ISP string `json:"ısp"`
	Organization string `json:"organization"`
}