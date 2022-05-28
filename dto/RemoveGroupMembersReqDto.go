package dto


type RemoveGroupMembersReqDto struct{
    UserIds  []string `json:"userIds"`
    Code  string `json:"code"`
}

