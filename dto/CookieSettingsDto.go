package dto


type CookieSettingsDto struct{
    CookieExpiresIn  int `json:"cookieExpiresIn"`
    CookieExpiresOnBrowserSession  bool `json:"cookieExpiresOnBrowserSession"`
}

