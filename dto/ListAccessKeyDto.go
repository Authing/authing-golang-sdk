package dto


type ListAccessKeyDto struct{
    UserId string `json:"userId,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
    Type string `json:"type,omitempty"`
    Status []string `json:"status,omitempty"`
}

