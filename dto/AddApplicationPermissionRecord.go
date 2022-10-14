package dto


type AddApplicationPermissionRecord struct{
    List  []ApplicationPermissionRecordItem `json:"list"`
    AppId  string `json:"appId"`
}

