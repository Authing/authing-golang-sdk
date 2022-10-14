package dto


type DeleteApplicationPermissionRecord struct{
    List  []DeleteApplicationPermissionRecordItem `json:"list"`
    AppId  string `json:"appId"`
}

