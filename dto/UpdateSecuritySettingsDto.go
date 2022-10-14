package dto


type UpdateSecuritySettingsDto struct{
    AllowedOrigins  []string `json:"allowedOrigins,omitempty"`
    AuthingTokenExpiresIn  int `json:"authingTokenExpiresIn,omitempty"`
    VerifyCodeLength  int `json:"verifyCodeLength,omitempty"`
    VerifyCodeMaxAttempts  int `json:"verifyCodeMaxAttempts,omitempty"`
    ChangeEmailStrategy  ChangeEmailStrategyDto `json:"changeEmailStrategy,omitempty"`
    ChangePhoneStrategy  ChangePhoneStrategyDto `json:"changePhoneStrategy,omitempty"`
    CookieSettings  CookieSettingsDto `json:"cookieSettings,omitempty"`
    RegisterDisabled  bool `json:"registerDisabled,omitempty"`
    RegisterAnomalyDetection  RegisterAnomalyDetectionConfigDto `json:"registerAnomalyDetection,omitempty"`
    CompletePasswordAfterPassCodeLogin  bool `json:"completePasswordAfterPassCodeLogin,omitempty"`
    LoginAnomalyDetection  LoginAnomalyDetectionConfigDto `json:"loginAnomalyDetection,omitempty"`
    LoginRequireEmailVerified  bool `json:"loginRequireEmailVerified,omitempty"`
    SelfUnlockAccount  SelfUnlockAccountConfigDto `json:"selfUnlockAccount,omitempty"`
    EnableLoginAccountSwitch  bool `json:"enableLoginAccountSwitch,omitempty"`
    QrcodeLoginStrategy  QrcodeLoginStrategyDto `json:"qrcodeLoginStrategy,omitempty"`
}

