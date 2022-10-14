package dto


type ListRoleDepartmentsDto struct{
    Code string `json:"code,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

