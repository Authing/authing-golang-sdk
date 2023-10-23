package dto


type UserSyncRelationDto struct{
    Provider  string `json:"provider"`
    UserIdInIdp  string `json:"userIdInIdp"`
    UserInfoInIdp  User `json:"userInfoInIdp"`
}

