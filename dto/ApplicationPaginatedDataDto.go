package dto


type ApplicationPaginatedDataDto struct{
    List  []ApplicationDto `json:"list"`
    TotalCount  int `json:"totalCount"`
}

