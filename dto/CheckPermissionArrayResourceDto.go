package dto

type CheckPermissionArrayResourceDto struct {
	Resources []string `json:"resources"`
	Action    string   `json:"action"`
}
