package dto


type UpdateMultipleTenantAdminDto struct{
    TenantIds  []string `json:"tenantIds"`
    ApiAuthorized  bool `json:"apiAuthorized,omitempty"`
}

