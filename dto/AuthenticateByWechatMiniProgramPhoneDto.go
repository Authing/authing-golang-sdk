package dto


type AuthenticateByWechatMiniProgramPhoneDto struct{
    EncryptedData  string `json:"encryptedData"`
    Iv  string `json:"iv"`
    Code  string `json:"code"`
}

