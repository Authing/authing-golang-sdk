package dto


type GetDataResourceDto struct{
    NamespaceCode string `json:"namespaceCode,omitempty"`
    ResourceCode string `json:"resourceCode,omitempty"`
}

