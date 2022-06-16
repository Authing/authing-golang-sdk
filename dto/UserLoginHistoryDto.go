package dto

type UserLoginHistoryDto struct {
	AppId       string `json:"appId"`
	AppName     string `json:"appName"`
	AppLogo     string `json:"appLogo"`
	AppLoginUrl string `json:"appLoginUrl"`
	ClientIp    string `json:"clientIp"`
	UserAgent   string `json:"userAgent,omitempty"`
	Time        string `json:"time"`
}
