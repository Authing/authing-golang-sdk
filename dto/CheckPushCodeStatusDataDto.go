package dto


type CheckPushCodeStatusDataDto struct{
    Status  string  `json:"status"`
    TokenSet  LoginTokenResponseDataDto `json:"tokenSet,omitempty"`
}

