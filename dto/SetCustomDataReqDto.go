package dto


type SetCustomDataReqDto struct{
    List  []SetCustomDataDto `json:"list"`
    TargetIdentifier  string `json:"targetIdentifier"`
    TargetType  string  `json:"targetType"`
    TenantId  string `json:"tenantId,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
}

