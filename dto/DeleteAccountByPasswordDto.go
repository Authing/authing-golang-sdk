package dto


type DeleteAccountByPasswordDto struct{
    Password  string `json:"password"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
}

