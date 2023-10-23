package dto


type ListApplicationAuthDto struct{
    AppId  string `json:"appId"`
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
    TargetName  string `json:"targetName,omitempty"`
    TargetTypeList  []string `json:"targetTypeList,omitempty"`
    Effect  string  `json:"effect,omitempty"`
    Enabled  bool `json:"enabled,omitempty"`
}

