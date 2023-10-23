package dto


type TenantListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []TenantRespDto `json:"list"`
}

