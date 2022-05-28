package dto


type ListResourcesDto struct{
    Namespace string `json:"namespace,omitempty"`
    Type string `json:"type,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

