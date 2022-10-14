package dto


type EmailTemplateDto struct{
    CustomizeEnabled  bool `json:"customizeEnabled"`
    Type  string  `json:"type"`
    Name  string `json:"name"`
    Subject  string `json:"subject"`
    Sender  string `json:"sender"`
    Content  string `json:"content"`
    ExpiresIn  int `json:"expiresIn,omitempty"`
    RedirectTo  string `json:"redirectTo,omitempty"`
    TplEngine  string  `json:"tplEngine,omitempty"`
}

