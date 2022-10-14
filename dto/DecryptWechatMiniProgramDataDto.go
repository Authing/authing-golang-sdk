package dto


type DecryptWechatMiniProgramDataDto struct{
    Code  string `json:"code"`
    Iv  string `json:"iv"`
    EncryptedData  string `json:"encryptedData"`
    ExtIdpConnidentifier  string `json:"extIdpConnidentifier"`
}

