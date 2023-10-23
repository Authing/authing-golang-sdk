package dto


type RemoveTenantUsersDto struct{
    TenantId  string `json:"tenantId"`
    LinkUserIds  []string `json:"linkUserIds,omitempty"`
    MemberIds  []string `json:"memberIds,omitempty"`
}

