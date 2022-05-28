package dto


type ListArchivedUsersSingleRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  ArchivedUsersListPagingDto `json:"data"`
}

