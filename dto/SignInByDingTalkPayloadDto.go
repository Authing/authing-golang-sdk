package dto


type SignInByDingTalkPayloadDto struct{
    Code  string `json:"code"`
    IsSnsCode  bool `json:"isSnsCode,omitempty"`
}

