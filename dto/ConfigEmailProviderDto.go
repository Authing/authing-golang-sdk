package dto


type ConfigEmailProviderDto struct{
    Type  string  `json:"type"`
    Enabled  bool `json:"enabled"`
    SmtpConfig  SMTPEmailProviderConfigInput `json:"smtpConfig,omitempty"`
    SendGridConfig  SendGridEmailProviderConfigInput `json:"sendGridConfig,omitempty"`
    AliExmailConfig  AliExmailEmailProviderConfigInput `json:"aliExmailConfig,omitempty"`
    TencentExmailConfig  TencentExmailEmailProviderConfigInput `json:"tencentExmailConfig,omitempty"`
}

