package dto


type AddTenantUsersDto struct{
    LinkUserIds  []string `json:"linkUserIds"`
    TenantId  string `json:"tenantId"`
}

