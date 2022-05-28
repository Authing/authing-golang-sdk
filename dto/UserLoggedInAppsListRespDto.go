package dto


type UserLoggedInAppsListRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  []UserLoggedInAppsDto `json:"data"`
}

