package dto


type GeneFastpassQRCodeInfoDto struct{
    Scene  string `json:"scene"`
    QrcodeId  string `json:"qrcodeId"`
    ApiHost  string `json:"apiHost"`
    User  FastpassUserInfoDto `json:"user"`
    AppId  string `json:"appId"`
    Userpool  FastpassUserInfoDto `json:"userpool"`
}

