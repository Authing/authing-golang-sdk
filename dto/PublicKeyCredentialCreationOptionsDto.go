package dto


type PublicKeyCredentialCreationOptionsDto struct{
    Challenge  string `json:"challenge"`
    ExcludeCredentials  []PublicKeyCredentialDescriptorDto `json:"excludeCredentials"`
    PubKeyCredParams  []PublicKeyCredentialParametersDto `json:"pubKeyCredParams"`
    Rp  PublicKeyCredentialRpEntityDto `json:"rp"`
    Timeout  int `json:"timeout"`
    User  PublicKeyCredentialUserEntityDto `json:"user"`
}

