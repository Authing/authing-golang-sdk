package dto


type GetExtIdpDto struct{
    Id string `json:"id,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
    AppId string `json:"appId,omitempty"`
    Type string `json:"type,omitempty"`
}

