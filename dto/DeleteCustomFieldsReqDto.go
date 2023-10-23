package dto


type DeleteCustomFieldsReqDto struct{
    TenantId  string `json:"tenantId"`
    List  []DeleteCustomFieldDto `json:"list"`
}

