package dto


type PreCheckCodeDto struct{
    CodeType  string  `json:"codeType"`
    SmsCodePayload  PreCheckSmsCodeDto `json:"smsCodePayload,omitempty"`
    EmailCodePayload  PreCheckEmailCodeDto `json:"emailCodePayload,omitempty"`
}

