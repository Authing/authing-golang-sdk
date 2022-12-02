package dto


type CheckParamsDataResourceRespDto struct{
    IsValid  bool `json:"isValid"`
    Message  string `json:"message,omitempty"`
}

