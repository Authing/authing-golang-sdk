package dto


type CreateExtIdpDto struct{
    Type  string  `json:"type"`
    Name  string `json:"name"`
    TenantId  string `json:"tenantId,omitempty"`
}

