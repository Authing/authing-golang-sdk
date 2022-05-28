package dto


type RoleAuthorizedResourcePaginatedRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  RoleAuthorizedResourcePagingDto `json:"data"`
}

