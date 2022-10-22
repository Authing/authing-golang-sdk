package dto


type GetResourceAuthorizedTargetDataDto struct{
    TotalCount  int `json:"totalCount"`
    List  []ResourceAuthorizedTargetDto `json:"list"`
}

