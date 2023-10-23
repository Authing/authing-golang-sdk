package dto


type GetUserAuthResourcePermissionList struct{
    NamespaceCode  string `json:"namespaceCode"`
    Actions  []string `json:"actions"`
    Resource  string `json:"resource"`
}

