package dto


type SetCustomFieldsReqDto struct{
    List  []SetCustomFieldDto `json:"list"`
    TenantId  string `json:"tenantId,omitempty"`
}

