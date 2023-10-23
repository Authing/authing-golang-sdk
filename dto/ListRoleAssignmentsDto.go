package dto


type ListRoleAssignmentsDto struct{
    RoleCode string `json:"roleCode,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Query string `json:"query,omitempty"`
    NamespaceCode string `json:"namespaceCode,omitempty"`
    TargetType []string `json:"targetType,omitempty"`
}

