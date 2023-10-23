package dto


type IpListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []IpListRespDto `json:"list"`
}

