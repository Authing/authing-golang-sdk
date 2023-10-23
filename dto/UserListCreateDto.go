package dto


type UserListCreateDto struct{
    ExpireAt  int `json:"expireAt"`
    LimitList  []string `json:"limitList"`
    RemoveType  string `json:"removeType"`
    AddType  string `json:"addType"`
    UserListType  string `json:"userListType"`
    UserIds  []string `json:"userIds"`
}

