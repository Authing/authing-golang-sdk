package dto


type ListTenantAdminDto struct{
    TenantId  string `json:"tenantId"`
    Keywords  string `json:"keywords,omitempty"`
    Page  string `json:"page,omitempty"`
    Limit  string `json:"limit,omitempty"`
}

