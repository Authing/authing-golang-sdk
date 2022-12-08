package dto


type GetAssignedAccountDto struct{
    AppId string `json:"appId,omitempty"`
    TargetType string `json:"targetType,omitempty"`
    TargetIdentifier string `json:"targetIdentifier,omitempty"`
}

