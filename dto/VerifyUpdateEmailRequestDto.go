package dto


type VerifyUpdateEmailRequestDto struct{
    EmailPasscodePayload  UpdateEmailByEmailPassCodeDto `json:"emailPasscodePayload"`
    VerifyMethod  string  `json:"verifyMethod"`
}

