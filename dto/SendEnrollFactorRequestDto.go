package dto


type SendEnrollFactorRequestDto struct{
    Profile  FactorProfile `json:"profile"`
    FactorType  string  `json:"factorType"`
}

