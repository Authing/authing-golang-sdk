package dto


type GetRoleAuthorizedResourcesDto struct{
    Code string `json:"code,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    ResourceType string `json:"resource_type,omitempty"`
}

