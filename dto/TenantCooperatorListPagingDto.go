package dto


type TenantCooperatorListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []TenantCooperatorDto `json:"list"`
}

