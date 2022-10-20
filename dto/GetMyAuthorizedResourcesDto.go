package dto


type GetMyAuthorizedResourcesDto struct{
    Namespace string `json:"namespace,omitempty"`
    ResourceType string `json:"resourceType,omitempty"`
}

