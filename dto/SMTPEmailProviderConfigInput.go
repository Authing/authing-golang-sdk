package dto


type SMTPEmailProviderConfigInput struct{
    SmtpHost  string `json:"smtpHost"`
    SmtpPort  int `json:"smtpPort"`
    Sender  string `json:"sender"`
    SenderPass  string `json:"senderPass"`
    EnableSSL  bool `json:"enableSSL"`
}

