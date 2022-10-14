package dto


type EmailProviderDto struct{
    Enabled  bool `json:"enabled"`
    Type  string  `json:"type,omitempty"`
    SmtpConfig  SMTPEmailProviderConfig `json:"smtpConfig,omitempty"`
    SendGridConfig  SendGridEmailProviderConfig `json:"sendGridConfig,omitempty"`
    AliExmailConfig  AliExmailEmailProviderConfig `json:"aliExmailConfig,omitempty"`
    TencentExmailConfig  TencentExmailEmailProviderConfig `json:"tencentExmailConfig,omitempty"`
}

