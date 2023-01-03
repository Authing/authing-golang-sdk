package dto

type CheckPermissionDto struct {
	Resources             []string      `json:"resources"`
	Action                string        `json:"action"`
	UserId                string        `json:"userId"`
	NamespaceCode         string        `json:"namespaceCode"`
	JudgeConditionEnabled bool          `json:"judgeConditionEnabled,omitempty"`
	AuthEnvParams         AuthEnvParams `json:"authEnvParams,omitempty"`
}
