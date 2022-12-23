package dto


type ListPermissionNamespacesDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Query string `json:"query,omitempty"`
}

