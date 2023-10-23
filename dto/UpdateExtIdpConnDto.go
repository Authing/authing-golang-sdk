package dto


type UpdateExtIdpConnDto struct{
    Id  string `json:"id"`
    DisplayName  string `json:"displayName"`
    Fields  interface{} `json:"fields"`
    Logo  string `json:"logo,omitempty"`
    LoginOnly  bool `json:"loginOnly,omitempty"`
    TenantId  string `json:"tenantId,omitempty"`
}

