package dto


type ListPermissionNamespaceRolesDto struct{
    Code string `json:"code,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Query string `json:"query,omitempty"`
}

