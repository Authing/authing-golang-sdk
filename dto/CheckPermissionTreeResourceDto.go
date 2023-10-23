package dto


type CheckPermissionTreeResourceDto struct{
    Resources  []string `json:"resources"`
    Action  string `json:"action"`
}

