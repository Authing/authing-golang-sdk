package dto


type ListAccessKeyRespDto struct{
    AccessKeyId  string `json:"accessKeyId"`
    AccessKeySecret  string `json:"accessKeySecret"`
    UserId  string `json:"userId"`
    CreatedAt  string `json:"createdAt"`
    Status  string `json:"status"`
    LastDate  string `json:"lastDate"`
    UserPoolId  string `json:"userPoolId"`
    Enable  bool `json:"enable"`
}

