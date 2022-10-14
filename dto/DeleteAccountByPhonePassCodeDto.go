package dto


type DeleteAccountByPhonePassCodeDto struct{
    PhoneNumber  string `json:"phoneNumber"`
    PassCode  string `json:"passCode"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
}

