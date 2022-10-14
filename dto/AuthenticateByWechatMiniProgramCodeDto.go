package dto


type AuthenticateByWechatMiniProgramCodeDto struct{
    EncryptedData  string `json:"encryptedData"`
    Iv  string `json:"iv"`
    Code  string `json:"code"`
}

