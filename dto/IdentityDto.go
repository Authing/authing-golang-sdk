package dto

type IdentityDto struct {
	IdentityId    string   `json:"identityId"`
	ExtIdpId      string   `json:"extIdpId"`
	Provider      string   `json:"provider"`
	Type          string   `json:"type"`
	UserIdInIdp   string   `json:"userIdInIdp"`
	UserInfoInIdp User     `json:"userInfoInIdp"`
	AccessToken   string   `json:"accessToken,omitempty"`
	RefreshToken  string   `json:"refreshToken,omitempty"`
	OriginConnIds []string `json:"originConnIds"`
}
