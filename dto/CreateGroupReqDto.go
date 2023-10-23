package dto


type CreateGroupReqDto struct{
    Type  string `json:"type"`
    Description  string `json:"description"`
    Name  string `json:"name"`
    Code  string `json:"code"`
    CustomData  interface{} `json:"customData,omitempty"`
}

