package dto


type SMTPEmailProviderConfig struct{
    SmtpHost  string `json:"smtp_host"`
    SmtpPort  int `json:"smtp_port"`
    Sender  string `json:"sender"`
    SenderPass  string `json:"senderPass"`
    Secure  bool `json:"secure"`
}

