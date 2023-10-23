package dto


type SignInByBaiduPayloadDto struct{
    Code  string `json:"code,omitempty"`
    AccessToken  string `json:"access_token,omitempty"`
}

