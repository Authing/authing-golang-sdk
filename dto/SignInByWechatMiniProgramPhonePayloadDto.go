package dto


type SignInByWechatMiniProgramPhonePayloadDto struct{
    EncryptedData  string `json:"encryptedData"`
    Iv  string `json:"iv"`
    Code  string `json:"code"`
}

