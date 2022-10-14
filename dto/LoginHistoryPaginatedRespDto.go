package dto


type LoginHistoryPaginatedRespDto struct{
    TotalCount  int `json:"totalCount"`
    List  []LoginHistoryDto `json:"list"`
}

