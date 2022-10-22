package dto


type UpdateGroupReqDto struct{
    Description  string `json:"description"`
    Code  string `json:"code"`
    Name  string `json:"name,omitempty"`
    NewCode  string `json:"newCode,omitempty"`
}

