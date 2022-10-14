package dto


type LoggedInAppsDto struct{
    AppId  string `json:"appId"`
    AppName  string `json:"appName"`
    AppLoginUrl  string `json:"appLoginUrl"`
    AppLogo  string `json:"appLogo"`
    Active  bool `json:"active"`
}

