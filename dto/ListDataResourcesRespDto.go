package dto


type ListDataResourcesRespDto struct{
    ResourceName  string `json:"resourceName"`
    ResourceCode  string `json:"resourceCode"`
    Type  string  `json:"type"`
    Description  string `json:"description,omitempty"`
    NamespaceCode  string `json:"namespaceCode"`
    NamespaceName  string `json:"namespaceName"`
    AuthorizationNum  int `json:"authorizationNum"`
    UpdatedAt  string `json:"updatedAt"`
}

