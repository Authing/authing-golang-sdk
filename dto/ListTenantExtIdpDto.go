package dto


type ListTenantExtIdpDto struct{
    TenantId string `json:"tenantId,omitempty"`
    AppId string `json:"appId,omitempty"`
    Type string `json:"type,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

