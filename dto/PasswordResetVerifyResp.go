package dto


type PasswordResetVerifyResp struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    RequestId  string `json:"requestId,omitempty"`
    Data  ResetPasswordVerify `json:"data"`
}

