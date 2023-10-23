package dto


type CreateExtIdpDto struct{
    Name  string `json:"name"`
    Type  string  `json:"type"`
    TenantId  string `json:"tenantId,omitempty"`
}

