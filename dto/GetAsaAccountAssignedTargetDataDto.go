package dto


type GetAsaAccountAssignedTargetDataDto struct{
    TotalCount  int `json:"totalCount"`
    List  []AsaAccountTargetDto `json:"list"`
}

