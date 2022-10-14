package dto


type ResetPasswordByPhonePassCodeDto struct{
    PhoneNumber  string `json:"phoneNumber"`
    PassCode  string `json:"passCode"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
}

