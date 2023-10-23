package dto


type TenantApplicationListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []TenantApplicationDto `json:"list"`
}

