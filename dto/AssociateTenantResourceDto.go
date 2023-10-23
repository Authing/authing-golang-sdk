package dto


type AssociateTenantResourceDto struct{
    Code  string `json:"code"`
    Association  bool `json:"association"`
    AppId  string `json:"appId"`
    TenantId  string `json:"tenantId,omitempty"`
}

