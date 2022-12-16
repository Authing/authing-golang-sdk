package dto


type CreateAccessKeyRespDto struct{
    AccessKeyId  string `json:"accessKeyId"`
    UserId  string `json:"userId"`
    CreatedAt  string `json:"createdAt"`
    Status  string `json:"status"`
    LastDate  string `json:"lastDate"`
    UserPoolId  string `json:"userPoolId"`
}

