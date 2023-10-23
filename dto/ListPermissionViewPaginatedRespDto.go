package dto


type ListPermissionViewPaginatedRespDto struct{
    TotalCount  int `json:"totalCount"`
    List  []ListPermissionViewRespDto `json:"list"`
}

