package dto


type CheckQRCodeStatusDataDto struct{
    Status  string  `json:"status"`
    Ticket  string `json:"ticket,omitempty"`
    BriefUserInfo  QRCodeStatusBriefUserInfoDto `json:"briefUserInfo,omitempty"`
    TokenSet  LoginTokenResponseDataDto `json:"tokenSet,omitempty"`
}

