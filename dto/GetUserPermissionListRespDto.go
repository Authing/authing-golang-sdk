package dto


type GetUserPermissionListRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  GetUserPermissionListDataDto `json:"data"`
}

