package dto


type GetSecurityInfoDto struct{
    PasswordSecurityLevel  int `json:"passwordSecurityLevel"`
    MfaEnrolled  bool `json:"mfaEnrolled"`
    PasswordSet  bool `json:"passwordSet"`
    EmailBinded  bool `json:"emailBinded"`
    PhoneBinded  bool `json:"phoneBinded"`
    SecurityScore  int `json:"securityScore"`
}

