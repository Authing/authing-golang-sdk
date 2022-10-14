package dto


type FactorDto struct{
    FactorId  string `json:"factorId"`
    FactorType  string  `json:"factorType"`
    Profile  interface{} `json:"profile"`
}

