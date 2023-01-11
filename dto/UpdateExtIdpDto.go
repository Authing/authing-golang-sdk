package dto

type UpdateExtIdpDto struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	TenantId string `json:"tenantId,omitempty"`
}
