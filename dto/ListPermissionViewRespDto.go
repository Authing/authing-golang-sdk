package dto


type ListPermissionViewRespDto struct{
    UserId  string `json:"userId"`
    Username  string `json:"username,omitempty"`
    NamespaceId  string `json:"namespaceId"`
    NamespaceCode  string `json:"namespaceCode"`
    NamespaceName  string `json:"namespaceName"`
    DataResourceId  string `json:"dataResourceId"`
    DataResourceCode  string `json:"dataResourceCode"`
    DataResourceName  string `json:"dataResourceName"`
    DataPolicyList  []PolicyBo `json:"dataPolicyList"`
    RoleList  []RoleBo `json:"roleList"`
    GroupList  []GroupBo `json:"groupList"`
    NodeList  []NodeBo `json:"nodeList"`
}

