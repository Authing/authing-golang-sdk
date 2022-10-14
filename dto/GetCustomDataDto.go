package dto


type GetCustomDataDto struct{
    TargetType string `json:"targetType,omitempty"`
    TargetIdentifier string `json:"targetIdentifier,omitempty"`
    Namespace string `json:"namespace,omitempty"`
}

