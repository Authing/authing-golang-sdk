package dto


type UpdateGroupReqDto struct{
    Description  string `json:"description"`
    Name  string `json:"name"`
    Code  string `json:"code"`
    NewCode  string `json:"newCode,omitempty"`
}

