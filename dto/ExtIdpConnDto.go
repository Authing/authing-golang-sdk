package dto

type ExtIdpConnDto struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Logo        string `json:"logo"`
	Identifier  string `json:"identifier,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
