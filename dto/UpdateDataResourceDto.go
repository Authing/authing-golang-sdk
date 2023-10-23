package dto


type UpdateDataResourceDto struct{
    ResourceCode  string `json:"resourceCode"`
    NamespaceCode  string `json:"namespaceCode"`
    ResourceName  string `json:"resourceName,omitempty"`
    Description  string `json:"description,omitempty"`
    Struct  interface{} `json:"struct,omitempty"`
    Actions  []string `json:"actions,omitempty"`
}

