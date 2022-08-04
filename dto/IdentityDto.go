package dto

type IdentityDto struct {
	IdentityId    string   `json:"identityId"`
	ExtIdpId      string   `json:"extIdpId"`
	Provider      string   `json:"provider"`
	Type          string   `json:"type"`
	UserIdInIdp   string   `json:"userIdInIdp"`
	OriginConnIds []string `json:"originConnIds"`
}
