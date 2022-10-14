package dto


type GenerateQrcodeDto struct{
    Type  string  `json:"type"`
    ExtIdpConnId  string `json:"extIdpConnId,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    Context  interface{} `json:"context,omitempty"`
    AutoMergeQrCode  bool `json:"autoMergeQrCode,omitempty"`
}

