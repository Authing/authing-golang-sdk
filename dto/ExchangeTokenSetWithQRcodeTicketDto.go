package dto


type ExchangeTokenSetWithQRcodeTicketDto struct{
    Ticket  string `json:"ticket"`
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
}

