package dto


type UpdateDataResourceDto struct{
    ResourceCode  string `json:"resourceCode"`
    NamespaceCode  string `json:"namespaceCode"`
    ResourceName  string `json:"resourceName,omitempty"`
    Description  string `json:"description,omitempty"`
    Struct  (DataResourceTreeStructs | string | []string) `json:"struct,omitempty"`
    Actions  []string `json:"actions,omitempty"`
}

