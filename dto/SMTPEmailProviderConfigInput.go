package dto


type SMTPEmailProviderConfigInput struct{
    SmtpHost  string `json:"smtp_host"`
    SmtpPort  int `json:"smtp_port"`
    Sender  string `json:"sender,omitempty"`
    SenderPass  string `json:"senderPass"`
    Secure  bool `json:"secure,omitempty"`
}

