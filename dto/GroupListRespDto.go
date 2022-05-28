package dto


type GroupListRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  []GroupDto `json:"data"`
}

