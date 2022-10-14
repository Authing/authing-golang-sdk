package dto


type GeneQRCodeDataDto struct{
    QrcodeId  string `json:"qrcodeId"`
    Url  string `json:"url"`
    CustomLogoUrl  string `json:"customLogoUrl,omitempty"`
}

