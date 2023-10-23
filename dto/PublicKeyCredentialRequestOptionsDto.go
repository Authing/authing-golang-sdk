package dto


type PublicKeyCredentialRequestOptionsDto struct{
    Challenge  string `json:"challenge"`
    AllowCredentials  []PublicKeyCredentialDescriptorDto `json:"allowCredentials,omitempty"`
    RpId  string `json:"rpId"`
    Timeout  int `json:"timeout"`
}

