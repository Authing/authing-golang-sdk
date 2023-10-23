package dto


type CheckExternalUserPermissionDto struct{
    Resources  []string `json:"resources"`
    Action  string `json:"action"`
    ExternalId  string `json:"externalId"`
    NamespaceCode  string `json:"namespaceCode"`
    JudgeConditionEnabled  bool `json:"judgeConditionEnabled,omitempty"`
    AuthEnvParams  AuthEnvParams `json:"authEnvParams,omitempty"`
}

