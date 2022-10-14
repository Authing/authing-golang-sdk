package dto


type SignInByMobileOptionsDto struct{
    Scope  string `json:"scope,omitempty"`
    Context  string `json:"context,omitempty"`
    TenantId  string `json:"tenantId,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

