package dto


type GetAuthorizedResourcesDto struct{
    TargetType string `json:"targetType,omitempty"`
    TargetIdentifier string `json:"targetIdentifier,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    ResourceType string `json:"resourceType,omitempty"`
    ResourceList []string `json:"resourceList,omitempty"`
    WithDenied bool `json:"withDenied,omitempty"`
}

