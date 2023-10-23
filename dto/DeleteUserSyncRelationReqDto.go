package dto


type DeleteUserSyncRelationReqDto struct{
    Provider  string `json:"provider"`
    UserId  string `json:"userId"`
    UserIdType  string  `json:"userIdType,omitempty"`
}

