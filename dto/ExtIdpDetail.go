package dto


type ExtIdpDetail struct{
    Id  string `json:"id"`
    Name  string `json:"name"`
    TenantId  string `json:"tenantId,omitempty"`
    Type  string `json:"type"`
    Connections  interface{} `json:"connections"`
    AutoJoin  bool `json:"autoJoin"`
}

