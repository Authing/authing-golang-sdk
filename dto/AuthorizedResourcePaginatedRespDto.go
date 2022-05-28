package dto


type AuthorizedResourcePaginatedRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  AuthorizedResourcePagingDto `json:"data"`
}

