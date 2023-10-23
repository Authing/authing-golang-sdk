package dto


type MultipleTenantAdminListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []MultipleTenantAdminDto `json:"list"`
}

