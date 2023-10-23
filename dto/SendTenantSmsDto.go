package dto


type SendTenantSmsDto struct{
    AdminName  string `json:"adminName"`
    UserName  string `json:"userName"`
    Identifier  string `json:"identifier"`
    Phone  string `json:"phone"`
    PhoneCountryCode  string `json:"phoneCountryCode"`
    TenantId  string `json:"tenantId"`
    TenantName  string `json:"tenantName"`
}

