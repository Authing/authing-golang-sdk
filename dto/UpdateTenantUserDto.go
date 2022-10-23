package dto


type UpdateTenantUserDto struct{
    Updates  interface{} `json:"updates"`
    TenantId  string `json:"tenantId"`
    LinkUserId  string `json:"linkUserId,omitempty"`
    MemberId  string `json:"memberId,omitempty"`
}

