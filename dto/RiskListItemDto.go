package dto


type RiskListItemDto struct{
    Id  string `json:"id"`
    UserId  string `json:"userId"`
    AddType  string `json:"addType"`
    UserListType  string `json:"userListType"`
    RemoveType  string `json:"removeType"`
    LimitList  []string `json:"limitList"`
    AddAt  string `json:"addAt"`
}

