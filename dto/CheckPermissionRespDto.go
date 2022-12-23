package dto


type CheckPermissionRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  CheckPermissionDataDto `json:"data"`
}

