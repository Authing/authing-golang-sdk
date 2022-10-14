package dto


type UserLoggedInIdentitiesRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    RequestId  string `json:"requestId,omitempty"`
    Data  []UserLoggedInIdentitiesDto `json:"data"`
}

