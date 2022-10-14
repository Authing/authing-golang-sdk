package dto


type VerifyUpdatePhoneRequestDto struct{
    PhonePassCodePayload  UpdatePhoneByPhonePassCodeDto `json:"phonePassCodePayload"`
    VerifyMethod  string  `json:"verifyMethod"`
}

