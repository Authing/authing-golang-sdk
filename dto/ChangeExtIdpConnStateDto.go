package dto


type ChangeExtIdpConnStateDto struct{
    AppId  string `json:"appId"`
    Enabled  bool `json:"enabled"`
    Id  string `json:"id"`
    TenantId  string `json:"tenantId,omitempty"`
    AppIds  []string `json:"appIds,omitempty"`
}

