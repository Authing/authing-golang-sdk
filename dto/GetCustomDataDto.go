package dto


type GetCustomDataDto struct{
    TargetType string `json:"target_type,omitempty"`
    TargetIdentifier string `json:"target_identifier,omitempty"`
    Namespace string `json:"namespace,omitempty"`
}

