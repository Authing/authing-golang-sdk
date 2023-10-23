package dto


type ListAuthSubjectDto struct{
    TargetType  string  `json:"targetType"`
    TargetId  string `json:"targetId"`
    AppName  string `json:"appName,omitempty"`
    AppTypeList  []string `json:"appTypeList,omitempty"`
    Effect  []string `json:"effect,omitempty"`
    Enabled  bool `json:"enabled,omitempty"`
}

