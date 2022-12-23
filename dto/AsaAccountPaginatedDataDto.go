package dto


type AsaAccountPaginatedDataDto struct{
    List  []AsaAccountDto `json:"list"`
    TotalCount  int `json:"totalCount"`
}

