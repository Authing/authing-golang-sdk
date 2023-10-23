package dto


type UserFieldDecryptReqDto struct{
    Data  []UserFieldDecryptReqItemDto `json:"data"`
    PrivateKey  string `json:"privateKey"`
}

