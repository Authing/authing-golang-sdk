package dto


type AuthenticateByPassCodeDto struct{
    PassCode  string `json:"passCode"`
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
}

