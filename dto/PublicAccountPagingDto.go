package dto


type PublicAccountPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []PublicAccountDto `json:"list"`
}

