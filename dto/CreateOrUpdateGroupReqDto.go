package dto


type CreateOrUpdateGroupReqDto struct{
    Type  string `json:"type"`
    Description  string `json:"description"`
    Name  string `json:"name"`
    Code  string `json:"code"`
}

