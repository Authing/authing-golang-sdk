package dto


type SetPublicAccountBatchReqDto struct{
    PublicAccountId  string `json:"publicAccountId"`
    UserIds  []string `json:"userIds"`
}

