package dto


type ListWebAuthnAuthenticatorDeviceDataDto struct{
    TotalCount  int `json:"totalCount"`
    List  []interface{} `json:"list"`
}

