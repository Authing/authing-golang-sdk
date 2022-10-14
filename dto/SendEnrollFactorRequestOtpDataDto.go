package dto


type SendEnrollFactorRequestOtpDataDto struct{
    QrCodeUri  string `json:"qrCodeUri"`
    QrCodeDataUrl  string `json:"qrCodeDataUrl"`
    RecoveryCode  string `json:"recoveryCode"`
}

