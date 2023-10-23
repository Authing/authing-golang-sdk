package dto


type GetSubjectAuthDataDto struct{
    AppId  string `json:"appId"`
    AppName  string `json:"appName"`
    ReqTargetId  string `json:"reqTargetId"`
    ReqTargetName  string `json:"reqTargetName"`
    ReqTargetType  string  `json:"reqTargetType"`
    TargetType  string  `json:"targetType"`
    TargetName  string `json:"targetName"`
    AuthType  string  `json:"authType"`
}

