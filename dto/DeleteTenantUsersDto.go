package dto


type DeleteTenantUsersDto struct{
    TenantId  string `json:"tenantId"`
    LinkUserIds  []string `json:"linkUserIds"`
    MemberIds  []string `json:"memberIds"`
}

