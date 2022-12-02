package dto


type DataResourcesPaginatedRespDto struct{
    TotalCount  int `json:"totalCount"`
    List  []ListDataResourcesRespDto `json:"list"`
}

