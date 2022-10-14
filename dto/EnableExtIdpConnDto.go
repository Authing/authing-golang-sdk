package dto


type EnableExtIdpConnDto struct{
    AppIds  string `json:"appIds"`
    AppId  string `json:"appId"`
    Enabled  bool `json:"enabled"`
    Id  string `json:"id"`
    TenantId  string `json:"tenantId,omitempty"`
}

