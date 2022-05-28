package dto


type UserLoginHistoryPaginatedRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  UserLoginHistoryPagingDto `json:"data"`
}

