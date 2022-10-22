package dto


type AuthorizeApplicationAccessDto struct{
    List  []ApplicationPermissionRecordItem `json:"list"`
    AppId  string `json:"appId"`
}

