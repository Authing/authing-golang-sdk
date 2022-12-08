package dto


type CreateArrayDataResourceDto struct{
    Actions  []string `json:"actions"`
    Struct  []string `json:"struct"`
    ResourceCode  string `json:"resourceCode"`
    ResourceName  string `json:"resourceName"`
    NamespaceCode  string `json:"namespaceCode"`
    Description  string `json:"description,omitempty"`
}

