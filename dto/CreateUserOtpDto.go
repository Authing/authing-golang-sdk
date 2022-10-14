package dto


type CreateUserOtpDto struct{
    Secret  string `json:"secret"`
    RecoveryCode  string `json:"recoveryCode,omitempty"`
}

