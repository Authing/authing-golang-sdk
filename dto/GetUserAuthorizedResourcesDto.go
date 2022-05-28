package dto


type GetUserAuthorizedResourcesDto struct{
    UserId string `json:"user_id,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    ResourceType string `json:"resource_type,omitempty"`
}

