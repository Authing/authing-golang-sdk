package dto


type SignupOptionsDto struct{
    ClientIp  string `json:"clientIp,omitempty"`
    PhonePassCodeForInformationCompletion  string `json:"phonePassCodeForInformationCompletion,omitempty"`
    EmailPassCodeForInformationCompletion  string `json:"emailPassCodeForInformationCompletion,omitempty"`
    Context  interface{} `json:"context"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
}

