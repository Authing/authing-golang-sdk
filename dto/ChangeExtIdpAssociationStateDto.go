package dto

type ChangeExtIdpAssociationStateDto struct {
	Id          string `json:"id"`
	Association bool   `json:"association"`
	TenantId    string `json:"tenantId,omitempty"`
}
