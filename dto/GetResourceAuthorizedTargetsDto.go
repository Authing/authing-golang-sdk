package dto


type GetResourceAuthorizedTargetsDto struct{
    Resource  string `json:"resource"`
    Namespace  string `json:"namespace,omitempty"`
    TargetType  string  `json:"targetType,omitempty"`
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
}

