package dto


type SetUserOfPublicAccountBatchReqDto struct{
    UserId  string `json:"userId"`
    PublicAccountIds  []string `json:"publicAccountIds"`
}

