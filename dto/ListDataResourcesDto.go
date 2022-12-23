package dto


type ListDataResourcesDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Query string `json:"query,omitempty"`
    NamespaceCodes string `json:"namespaceCodes,omitempty"`
}

