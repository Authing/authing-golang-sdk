package dto


type CreateTreeDataResourceDto struct{
    Actions  []string `json:"actions"`
    Struct  []DataResourceTreeStructs `json:"struct"`
    ResourceCode  string `json:"resourceCode"`
    ResourceName  string `json:"resourceName"`
    NamespaceCode  string `json:"namespaceCode"`
    Description  string `json:"description,omitempty"`
}

