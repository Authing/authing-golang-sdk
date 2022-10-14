package dto


type BindPhoneDto struct{
    PassCode  string `json:"passCode"`
    PhoneNumber  string `json:"phoneNumber"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
}

