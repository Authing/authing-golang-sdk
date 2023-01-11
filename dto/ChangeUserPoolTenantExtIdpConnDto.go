package dto

type ChangeUserPoolTenantExtIdpConnDto struct {
	Enabled bool     `json:"enabled"`
	ConnIds []string `json:"connIds"`
}
