package dto


type ResignUserReqDto struct{
    UserId  string `json:"userId"`
    UserIdType  string  `json:"userIdType,omitempty"`
}

