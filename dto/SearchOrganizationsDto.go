package dto

type SearchOrganizationsDto struct {
	Keywords       string `json:"keywords,omitempty"`
	Page           int    `json:"page,omitempty"`
	Limit          int    `json:"limit,omitempty"`
	WithCustomData bool   `json:"withCustomData,omitempty"`
	TenantId       string `json:"tenantId,omitempty"`
}
