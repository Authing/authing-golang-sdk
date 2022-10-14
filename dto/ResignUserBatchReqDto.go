package dto


type ResignUserBatchReqDto struct{
    UserIds  []string `json:"userIds"`
    UserIdType  string  `json:"userIdType,omitempty"`
}

