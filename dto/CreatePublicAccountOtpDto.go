package dto


type CreatePublicAccountOtpDto struct{
    Secret  string `json:"secret"`
    RecoveryCode  string `json:"recoveryCode,omitempty"`
}

