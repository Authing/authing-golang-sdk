package dto


type GetAuthorizedTargetDataDto struct{
    TotalCount  int `json:"totalCount"`
    List  []ResourcePermissionAssignmentDto `json:"list"`
}

