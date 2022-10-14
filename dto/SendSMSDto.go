package dto


type SendSMSDto struct{
    Channel  string  `json:"channel"`
    PhoneNumber  string `json:"phoneNumber"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
}

