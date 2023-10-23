package dto


type BindByPhoneCodeInputApi struct{
    Key  string `json:"key"`
    Action  string `json:"action"`
    Code  string `json:"code"`
    Phone  string `json:"phone"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
}

