package dto

type LoginIpWhitelistCheckConfigDto struct {
	Enabled     bool   `json:"enabled"`
	IpWhitelist string `json:"ipWhitelist"`
}
