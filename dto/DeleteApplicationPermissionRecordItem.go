package dto


type DeleteApplicationPermissionRecordItem struct{
    TargetType  string  `json:"targetType"`
    NamespaceCode  string `json:"namespaceCode,omitempty"`
    TargetIdentifier  []string `json:"targetIdentifier"`
}

