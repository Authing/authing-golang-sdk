package dto


type ImportOtpItemDataDto struct{
    Secret  string `json:"secret"`
    RecoveryCode  string `json:"recoveryCode,omitempty"`
}

