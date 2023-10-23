package dto


type ListApplicationAuthDataDto struct{
    Id  string `json:"id"`
    TargetId  string `json:"targetId"`
    TargetName  string `json:"targetName"`
    TargetType  string  `json:"targetType"`
    Effect  string  `json:"effect"`
    Enabled  bool `json:"enabled"`
    PermissionType  string  `json:"permissionType"`
}

