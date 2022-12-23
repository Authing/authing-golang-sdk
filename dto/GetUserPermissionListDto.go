package dto


type GetUserPermissionListDto struct{
    UserIds  []string `json:"userIds"`
    NamespaceCodes  []string `json:"namespaceCodes,omitempty"`
}

