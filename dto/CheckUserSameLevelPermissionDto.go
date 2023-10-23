package dto


type CheckUserSameLevelPermissionDto struct{
    ResourceNodeCodes  []string `json:"resourceNodeCodes"`
    Resource  string `json:"resource"`
    Action  string `json:"action"`
    UserId  string `json:"userId"`
    NamespaceCode  string `json:"namespaceCode"`
    JudgeConditionEnabled  bool `json:"judgeConditionEnabled,omitempty"`
    AuthEnvParams  AuthEnvParams `json:"authEnvParams,omitempty"`
}

