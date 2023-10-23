package dto


type SetCustomDataDto struct{
    TenantId  string `json:"tenantId,omitempty"`
    Key  string `json:"key"`
    Value  string `json:"value"`
}

