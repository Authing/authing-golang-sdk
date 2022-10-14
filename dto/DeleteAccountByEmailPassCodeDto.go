package dto


type DeleteAccountByEmailPassCodeDto struct{
    Email  string `json:"email,omitempty"`
    PassCode  string `json:"passCode"`
}

