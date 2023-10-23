package dto


type InviteTenantUsersDto struct{
    ErrMsgs  []errorEmailMsg `json:"errMsgs"`
    List  []InviteTenantUserRecord `json:"list"`
}

