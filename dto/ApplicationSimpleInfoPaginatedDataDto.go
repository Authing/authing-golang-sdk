package dto


type ApplicationSimpleInfoPaginatedDataDto struct{
    List  []ApplicationSimpleInfoDto `json:"list"`
    TotalCount  int `json:"totalCount"`
}

