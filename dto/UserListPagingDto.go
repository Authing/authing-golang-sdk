package dto


type UserListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []RiskListItemDto `json:"list"`
}

