package dto


type ResourceAuthorizedTargetDto struct{
    TargetType  string  `json:"targetType"`
    TargetIdentifier  string `json:"targetIdentifier"`
    Actions  []string `json:"actions"`
}

