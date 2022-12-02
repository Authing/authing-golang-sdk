package dto


type CommonResourceDto struct{
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
    Name  string `json:"name,omitempty"`
    Type  string  `json:"type"`
    Actions  []ResourceAction `json:"actions,omitempty"`
    ApiIdentifier  string `json:"apiIdentifier,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
    LinkedToTenant  bool `json:"linkedToTenant,omitempty"`
    Id  string `json:"id"`
    NamespaceId  int `json:"namespaceId"`
    NamespaceName  string `json:"namespaceName"`
    UserPoolId  string `json:"userPoolId"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
}

