package dto


type ListCommonResourceDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Keyword string `json:"keyword,omitempty"`
    NamespaceCodeList string `json:"namespaceCodeList,omitempty"`
}

