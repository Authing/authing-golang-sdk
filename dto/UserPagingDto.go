package dto


type UserPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []UserDto `json:"list"`
}

