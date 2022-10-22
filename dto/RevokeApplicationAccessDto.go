package dto


type RevokeApplicationAccessDto struct{
    List  []DeleteApplicationPermissionRecordItem `json:"list"`
    AppId  string `json:"appId"`
}

