package dto


type ResetPasswordByEmailPassCodeDto struct{
    Email  string `json:"email,omitempty"`
    PassCode  string `json:"passCode"`
}

