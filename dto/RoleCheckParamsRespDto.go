package dto


type RoleCheckParamsRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  CheckRoleParamsRespDto `json:"data"`
}

