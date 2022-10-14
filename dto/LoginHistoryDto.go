package dto


type LoginHistoryDto struct{
    UserId  string `json:"userId"`
    AppId  string `json:"appId"`
    AppName  string `json:"appName"`
    AppLoginUrl  string `json:"appLoginUrl"`
    AppLogo  string `json:"appLogo"`
    LoginAt  string `json:"loginAt"`
    ClientIp  string `json:"clientIp"`
    Success  bool `json:"success"`
    ErrorMessage  string `json:"errorMessage,omitempty"`
    UserAgent  string `json:"userAgent"`
    ParsedUserAgent  ParsedUserAgent `json:"parsedUserAgent"`
    LoginMethod  string `json:"loginMethod"`
    Geoip  GeoIp `json:"geoip"`
}

