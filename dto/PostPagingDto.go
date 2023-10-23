package dto


type PostPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []PostItemDto `json:"list"`
}

