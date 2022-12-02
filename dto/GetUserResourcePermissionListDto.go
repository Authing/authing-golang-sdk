package dto


type GetUserResourcePermissionListDto struct{
    Resources  []string `json:"resources"`
    UserId  string `json:"userId"`
    NamespaceCode  string `json:"namespaceCode"`
}

