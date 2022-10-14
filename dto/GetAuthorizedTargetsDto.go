package dto


type GetAuthorizedTargetsDto struct{
    Resource  string `json:"resource"`
    Namespace  string `json:"namespace,omitempty"`
    ResourceType  string  `json:"resourceType,omitempty"`
    TargetType  string  `json:"targetType,omitempty"`
    Actions  GetAuthorizedResourceActionDto `json:"actions,omitempty"`
}

