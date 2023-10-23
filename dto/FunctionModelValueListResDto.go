package dto


type FunctionModelValueListResDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  FunctionModelValueListDto `json:"data"`
}

