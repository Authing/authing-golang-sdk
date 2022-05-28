package dto


type OrganizationPaginatedRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  OrganizationPagingDto `json:"data"`
}

