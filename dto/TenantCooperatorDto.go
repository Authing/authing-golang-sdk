package dto


type TenantCooperatorDto struct{
    UserId  string `json:"userId"`
    Type  string `json:"type"`
    External  bool `json:"external"`
    User  UserDto `json:"user"`
    TenantUser  TenantUserDto `json:"tenantUser,omitempty"`
}

