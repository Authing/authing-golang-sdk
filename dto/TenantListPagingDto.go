package dto


type TenantListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []UpdateTenantDto `json:"list"`
}

