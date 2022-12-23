package dto


type ListRolesDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Keywords string `json:"keywords,omitempty"`
    Namespace string `json:"namespace,omitempty"`
}

