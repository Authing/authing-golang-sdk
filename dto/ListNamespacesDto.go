package dto


type ListNamespacesDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Keywords string `json:"keywords,omitempty"`
}

