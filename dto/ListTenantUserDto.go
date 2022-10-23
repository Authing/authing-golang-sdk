package dto


type ListTenantUserDto struct{
    TenantId  string `json:"tenantId"`
    Keywords  string `json:"keywords,omitempty"`
    Page  string `json:"page,omitempty"`
    Limit  string `json:"limit,omitempty"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
    WithIdentities  bool `json:"withIdentities,omitempty"`
    WithDepartmentIds  bool `json:"withDepartmentIds,omitempty"`
}

