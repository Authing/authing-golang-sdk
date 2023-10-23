package dto


type EventAppPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []EventAppDto `json:"list"`
}

