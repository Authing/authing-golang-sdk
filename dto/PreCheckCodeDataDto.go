package dto


type PreCheckCodeDataDto struct{
    IsValid  bool `json:"isValid"`
    Message  string `json:"message,omitempty"`
}

