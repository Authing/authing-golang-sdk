package dto


type ImportOtpItemDto struct{
    UserId  string `json:"userId"`
    Otp  ImportOtpItemDataDto `json:"otp"`
}

