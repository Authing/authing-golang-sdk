package dto


type GetTenantDto struct{
    TenantId string `json:"tenantId,omitempty"`
    WithMembersCount bool `json:"withMembersCount,omitempty"`
    WithAppDetail bool `json:"withAppDetail,omitempty"`
}

