package dto


type CommonResourcePagingDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    RequestId  string `json:"requestId,omitempty"`
    TotalCount  int `json:"totalCount"`
    List  []CommonResourceDto `json:"list"`
}

