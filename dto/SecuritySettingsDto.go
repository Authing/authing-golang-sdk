package dto


type SecuritySettingsDto struct{
    AllowedOrigins  []string `json:"allowedOrigins,omitempty"`
    AuthingTokenExpiresIn  int `json:"authingTokenExpiresIn"`
    VerifyCodeLength  int `json:"verifyCodeLength"`
    VerifyCodeMaxAttempts  int `json:"verifyCodeMaxAttempts"`
    ChangeEmailStrategy  ChangeEmailStrategyDto `json:"changeEmailStrategy"`
    ChangePhoneStrategy  ChangePhoneStrategyDto `json:"changePhoneStrategy"`
    CookieSettings  CookieSettingsDto `json:"cookieSettings,omitempty"`
    RegisterDisabled  bool `json:"registerDisabled"`
    RegisterAnomalyDetection  RegisterAnomalyDetectionConfigDto `json:"registerAnomalyDetection"`
    CompletePasswordAfterPassCodeLogin  bool `json:"completePasswordAfterPassCodeLogin"`
    LoginAnomalyDetection  LoginAnomalyDetectionConfigDto `json:"loginAnomalyDetection"`
    LoginRequireEmailVerified  bool `json:"loginRequireEmailVerified"`
    SelfUnlockAccount  SelfUnlockAccountConfigDto `json:"selfUnlockAccount"`
    EnableLoginAccountSwitch  bool `json:"enableLoginAccountSwitch"`
    QrcodeLoginStrategy  QrcodeLoginStrategyDto `json:"qrcodeLoginStrategy"`
}

