package dto

type AuthEnvParams struct {
	Ip          string `json:"ip,omitempty"`
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
	Country     string `json:"country,omitempty"`
	DeviceType  string `json:"deviceType,omitempty"`
	SystemType  string `json:"systemType,omitempty"`
	BrowserType string `json:"browserType,omitempty"`
	RequestDate string `json:"requestDate,omitempty"`
}
