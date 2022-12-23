package dto


type ListNamespaceRolesDto struct{
    Code string `json:"code,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Keywords string `json:"keywords,omitempty"`
}

