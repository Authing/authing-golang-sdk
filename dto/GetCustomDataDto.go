package dto


type GetCustomDataDto struct{
    TenantId string `json:"tenantId,omitempty"`
    TargetType string `json:"targetType,omitempty"`
    TargetIdentifier string `json:"targetIdentifier,omitempty"`
    Namespace string `json:"namespace,omitempty"`
}

