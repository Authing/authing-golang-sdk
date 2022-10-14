package dto


type GetUserAuthorizedResourcesDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    ResourceType string `json:"resourceType,omitempty"`
}

