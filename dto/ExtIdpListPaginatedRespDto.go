package dto


type ExtIdpListPaginatedRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  ExtIdpListPagingDto `json:"data"`
}

