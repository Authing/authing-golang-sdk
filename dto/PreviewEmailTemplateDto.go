package dto


type PreviewEmailTemplateDto struct{
    Type  string  `json:"type"`
    Content  string `json:"content,omitempty"`
    Subject  string `json:"subject,omitempty"`
    Sender  string `json:"sender,omitempty"`
    ExpiresIn  int `json:"expiresIn,omitempty"`
    TplEngine  string  `json:"tplEngine,omitempty"`
}

