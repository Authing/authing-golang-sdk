package dto


type ApplicationAgreementDto struct{
    DisplayAt  []string `json:"displayAt"`
    IsRequired  bool `json:"isRequired"`
    Lang  string  `json:"lang"`
    Content  string `json:"content"`
}

