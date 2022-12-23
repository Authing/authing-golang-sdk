package dto

type CheckPermissionStringResourceDto struct {
	Resources []string `json:"resources"`
	Action    string   `json:"action"`
}
