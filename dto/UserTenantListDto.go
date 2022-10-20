package dto


type UserTenantListDto struct{
    TenantId  string `json:"tenantId"`
    TenantName  string `json:"tenantName"`
    JoinAt  string `json:"joinAt"`
}

