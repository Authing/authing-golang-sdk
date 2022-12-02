package dto


type UserPermissionListDto struct{
    UserId  string `json:"userId"`
    NamespaceCode  string `json:"namespaceCode"`
    ResourceList  []OpenResource `json:"resourceList,omitempty"`
}

