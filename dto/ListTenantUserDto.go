package dto


type ListTenantUserDto struct{
    TenantId  string `json:"tenantId"`
    Keywords  string `json:"keywords,omitempty"`
    Options  ListTenantUsersOptionsDto `json:"options,omitempty"`
}

