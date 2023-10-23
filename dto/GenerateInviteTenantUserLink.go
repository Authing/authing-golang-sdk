package dto


type GenerateInviteTenantUserLink struct{
    ValidityTerm  string `json:"validityTerm"`
    Emails  []string `json:"emails"`
    AppId  string `json:"appId"`
    TenantId  string `json:"tenantId,omitempty"`
}

