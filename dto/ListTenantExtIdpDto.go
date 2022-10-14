package dto


type ListTenantExtIdpDto struct{
    TenantId string `json:"tenantId,omitempty"`
    AppId string `json:"appId,omitempty"`
    Type string `json:"type,omitempty"`
    Page string `json:"page,omitempty"`
    Limit string `json:"limit,omitempty"`
}

