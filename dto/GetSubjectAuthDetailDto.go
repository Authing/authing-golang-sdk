package dto


type GetSubjectAuthDetailDto struct{
    TargetId string `json:"targetId,omitempty"`
    TargetType string `json:"targetType,omitempty"`
    AppId string `json:"appId,omitempty"`
}

