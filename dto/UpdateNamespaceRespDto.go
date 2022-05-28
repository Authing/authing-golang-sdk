package dto


type UpdateNamespaceRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  UpdateNamespaceDto `json:"data"`
}

