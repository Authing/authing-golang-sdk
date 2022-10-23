package dto


type GetTenantUserDto struct{
    TenantId string `json:"tenantId,omitempty"`
    LinkUserId string `json:"linkUserId,omitempty"`
    MemberId string `json:"memberId,omitempty"`
}

