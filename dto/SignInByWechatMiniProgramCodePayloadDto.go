package dto


type SignInByWechatMiniProgramCodePayloadDto struct{
    EncryptedData  string `json:"encryptedData,omitempty"`
    Iv  string `json:"iv,omitempty"`
    Code  string `json:"code"`
}

