package dto


type ExtIdpDto struct{
    Id  string `json:"id"`
    Name  string `json:"name"`
    Logo  string `json:"logo"`
    TenantId  string `json:"tenantId,omitempty"`
    Type  string `json:"type"`
}

