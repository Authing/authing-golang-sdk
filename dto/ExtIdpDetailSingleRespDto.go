package dto


type ExtIdpDetailSingleRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  ExtIdpDetail `json:"data"`
}

