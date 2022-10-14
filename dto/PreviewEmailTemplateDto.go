package dto


type PreviewEmailTemplateDto struct{
    Sender  string `json:"sender"`
    Type  string  `json:"type"`
    Content  string `json:"content,omitempty"`
    Subject  string `json:"subject,omitempty"`
    ExpiresIn  int `json:"expiresIn,omitempty"`
    TplEngine  string  `json:"tplEngine,omitempty"`
}

