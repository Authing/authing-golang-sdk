package dto


type PreCheckSmsCodeDto struct{
    Phone  string `json:"phone"`
    PassCode  string `json:"passCode"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    Channel  string  `json:"channel"`
}

