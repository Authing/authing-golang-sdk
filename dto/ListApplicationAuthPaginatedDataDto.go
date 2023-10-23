package dto


type ListApplicationAuthPaginatedDataDto struct{
    List  []ListApplicationAuthDataDto `json:"list"`
    TotalCount  int `json:"totalCount"`
}

