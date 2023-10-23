package dto


type ListTenantApplicationsDto struct{
    Page string `json:"page,omitempty"`
    Limit string `json:"limit,omitempty"`
    Keywords string `json:"keywords,omitempty"`
    SsoEnabled bool `json:"sso_enabled,omitempty"`
}

