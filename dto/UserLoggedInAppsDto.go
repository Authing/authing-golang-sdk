package dto

type UserLoggedInAppsDto struct {
	AppId       string `json:"appId"`
	AppName     string `json:"appName"`
	AppLogo     string `json:"appLogo"`
	AppLoginUrl string `json:"appLoginUrl"`
}
