package dto


type ApplicationSimpleInfoDto struct{
    AppId  string `json:"appId"`
    AppIdentifier  string `json:"appIdentifier"`
    AppName  string `json:"appName"`
    AppLogo  string `json:"appLogo"`
    AppDescription  string `json:"appDescription,omitempty"`
    AppType  string  `json:"appType"`
    IsIntegrateApp  bool `json:"isIntegrateApp"`
}

