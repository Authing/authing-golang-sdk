package dto


type CreateAccessKeyDto struct{
    Type  string `json:"type"`
    UserId  string `json:"userId,omitempty"`
    TenantId  string `json:"tenantId,omitempty"`
}

