package dto


type FunctionModelValueResDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  RowDto `json:"data"`
}

