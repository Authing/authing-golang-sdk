package dto

type DeleteRoleDto struct {
	CodeList  []string `json:"codeList"`
	Namespace string   `json:"namespace,omitempty"`
}
