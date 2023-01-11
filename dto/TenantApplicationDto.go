package dto

type TenantApplicationDto struct {
	UserPoolId      string `json:"userPoolId"`
	TenantAppId     string `json:"tenantAppId"`
	Name            string `json:"name"`
	Description     string `json:"description,omitempty"`
	Logo            string `json:"logo"`
	ApplicationType string `json:"applicationType"`
	SsoEnabled      bool   `json:"ssoEnabled"`
	AppId           string `json:"appId"`
}
