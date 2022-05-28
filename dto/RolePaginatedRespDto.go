package dto


type RolePaginatedRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  RolePagingDto `json:"data"`
}

