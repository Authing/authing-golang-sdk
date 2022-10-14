package dto


type AppDto struct{
    AppId  string `json:"appId"`
    AppName  string `json:"appName"`
    AppLogo  string `json:"appLogo"`
    AppLoginUrl  string `json:"appLoginUrl"`
    AppDefaultLoginStrategy  string  `json:"appDefaultLoginStrategy"`
}

