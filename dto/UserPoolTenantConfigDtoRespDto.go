package dto


type UserPoolTenantConfigDtoRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    Data  UserPoolTenantConfigDto `json:"data"`
}

