package dto


type AuthorizedResourcePagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []AuthorizedResourceDto `json:"list"`
}

