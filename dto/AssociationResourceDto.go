package dto


type AssociationResourceDto struct{
    AppId  string `json:"appId"`
    Association  bool `json:"association"`
    Code  string `json:"code"`
    TenantId  string `json:"tenantId,omitempty"`
}

