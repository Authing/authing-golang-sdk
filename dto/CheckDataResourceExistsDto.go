package dto


type CheckDataResourceExistsDto struct{
    NamespaceCode string `json:"namespaceCode,omitempty"`
    ResourceName string `json:"resourceName,omitempty"`
    ResourceCode string `json:"resourceCode,omitempty"`
}

