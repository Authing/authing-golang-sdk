package dto


type RevokeApplicationAccessDto struct{
    AppId  string `json:"appId"`
    List  []DeleteApplicationPermissionRecordItem `json:"list"`
}

