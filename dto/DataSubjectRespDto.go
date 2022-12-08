package dto


type DataSubjectRespDto struct{
    TargetIdentifier  string `json:"targetIdentifier"`
    TargetType  string  `json:"targetType"`
    TargetName  string `json:"targetName"`
    AuthorizationTime  string `json:"authorizationTime"`
}

