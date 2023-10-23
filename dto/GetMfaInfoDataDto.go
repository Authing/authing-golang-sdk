package dto


type GetMfaInfoDataDto struct{
    MfaToken  string `json:"mfaToken"`
    MfaPhone  string `json:"mfaPhone,omitempty"`
    MfaPhoneCountryCode  string `json:"mfaPhoneCountryCode,omitempty"`
    MfaEmail  string `json:"mfaEmail,omitempty"`
    Nickname  string `json:"nickname,omitempty"`
    Username  string `json:"username,omitempty"`
    Phone  string `json:"phone,omitempty"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    FaceMfaEnabled  bool `json:"faceMfaEnabled,omitempty"`
    TotpMfaEnabled  bool `json:"totpMfaEnabled,omitempty"`
    ApplicationMfa  []ApplicationMfaDto `json:"applicationMfa"`
}

