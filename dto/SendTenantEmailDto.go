package dto


type SendTenantEmailDto struct{
    AdminName  string `json:"adminName"`
    UserName  string `json:"userName"`
    Email  string `json:"email"`
    Identifier  string `json:"identifier"`
    TenantId  string `json:"tenantId"`
    TenantName  string `json:"tenantName"`
}

