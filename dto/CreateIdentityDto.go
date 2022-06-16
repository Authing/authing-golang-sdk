package dto

type CreateIdentityDto struct {
	ExtIdpId    string `json:"extIdpId"`
	Provider    string `json:"provider"`
	Type        string `json:"type"`
	UserIdInIdp string `json:"userIdInIdp"`
}
