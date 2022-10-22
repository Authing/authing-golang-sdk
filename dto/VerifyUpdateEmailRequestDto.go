package dto


type VerifyUpdateEmailRequestDto struct{
    EmailPassCodePayload  UpdateEmailByEmailPassCodeDto `json:"emailPassCodePayload"`
    VerifyMethod  string  `json:"verifyMethod"`
}

