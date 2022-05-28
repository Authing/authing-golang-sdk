package dto


type ExtIdpConnDetailSingleRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  ExtIdpConnDetail `json:"data"`
}

