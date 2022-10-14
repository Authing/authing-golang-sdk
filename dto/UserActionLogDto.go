package dto


type UserActionLogDto struct{
    UserId  string `json:"userId"`
    UserAvatar  string `json:"userAvatar"`
    UserDisplayName  string `json:"userDisplayName"`
    UserLoginsCount  int `json:"userLoginsCount"`
    AppId  string `json:"appId"`
    AppName  string `json:"appName"`
    ClientIp  string `json:"clientIp,omitempty"`
    EventType  string  `json:"eventType"`
    EventDetail  string `json:"eventDetail,omitempty"`
    Success  bool `json:"success"`
    AppLoginUrl  string `json:"appLoginUrl"`
    AppLogo  string `json:"appLogo"`
    UserAgent  string `json:"userAgent"`
    ParsedUserAgent  ParsedUserAgent `json:"parsedUserAgent"`
    Geoip  GeoIp `json:"geoip"`
    Timestamp  string `json:"timestamp"`
    RequestId  string `json:"requestId"`
}

