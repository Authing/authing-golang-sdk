package dto


type InviteTenantUserRecordListDto struct{
    TotalCount  int `json:"totalCount"`
    List  []InviteTenantUserRecord `json:"list"`
}

