package dto


type UpdateEmailTemplateDto struct{
    Content  string `json:"content"`
    Sender  string `json:"sender"`
    Subject  string `json:"subject"`
    Name  string `json:"name"`
    CustomizeEnabled  bool `json:"customizeEnabled"`
    Type  string  `json:"type"`
    ExpiresIn  int `json:"expiresIn,omitempty"`
    RedirectTo  string `json:"redirectTo,omitempty"`
    TplEngine  string  `json:"tplEngine,omitempty"`
}

