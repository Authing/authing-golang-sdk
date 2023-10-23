package dto


type SetUserGroupsDto struct{
    GroupIds  []string `json:"groupIds"`
    UserId  string `json:"userId"`
}

