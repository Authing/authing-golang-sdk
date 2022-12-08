package dto


type UserMfaRespDto struct{
    TotpStatus  string `json:"totpStatus"`
    FaceMfaStatus  string `json:"faceMfaStatus"`
    SmsMfaStatus  string `json:"smsMfaStatus"`
    EmailMfaStatus  string `json:"emailMfaStatus"`
}

