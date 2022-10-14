package dto


type QrcodeLoginStrategyDto struct{
    QrcodeExpiresIn  int `json:"qrcodeExpiresIn"`
    TicketExpiresIn  int `json:"ticketExpiresIn"`
    AllowExchangeUserInfoFromBrowser  bool `json:"allowExchangeUserInfoFromBrowser"`
    ReturnFullUserInfo  bool `json:"returnFullUserInfo"`
}

