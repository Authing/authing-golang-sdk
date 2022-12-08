package dto


type QrcodeLoginStrategyDto struct{
    QrcodeExpiresIn  int `json:"qrcodeExpiresIn"`
    QrcodeExpiresInUnit  string `json:"qrcodeExpiresInUnit,omitempty"`
    TicketExpiresIn  int `json:"ticketExpiresIn"`
    TicketExpiresInUnit  string `json:"ticketExpiresInUnit,omitempty"`
    AllowExchangeUserInfoFromBrowser  bool `json:"allowExchangeUserInfoFromBrowser"`
    ReturnFullUserInfo  bool `json:"returnFullUserInfo"`
}

