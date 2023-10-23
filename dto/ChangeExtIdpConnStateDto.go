package dto


type ChangeExtIdpConnStateDto struct{
    Id  string `json:"id"`
    Enabled  bool `json:"enabled"`
    AppId  string `json:"appId"`
    TenantId  string `json:"tenantId,omitempty"`
    AppIds  []string `json:"appIds,omitempty"`
}

