package dto


type ListRolesDto struct{
    Namespace string `json:"namespace,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

