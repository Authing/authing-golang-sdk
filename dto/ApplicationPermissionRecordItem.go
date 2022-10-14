package dto


type ApplicationPermissionRecordItem struct{
    TargetType  string  `json:"targetType"`
    NamespaceCode  string `json:"namespaceCode,omitempty"`
    InheritByChildren  bool `json:"inheritByChildren,omitempty"`
    TargetIdentifier  []string `json:"targetIdentifier"`
    Effect  string  `json:"effect"`
}

