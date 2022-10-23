package dto


type TenantUserListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []TenantUserDto `json:"list"`
}

