package dto

type GetUserResourceStructDataDto struct {
	NamespaceCode string            `json:"namespaceCode"`
	ResourceCode  string            `json:"resourceCode"`
	PermissionBo  TreePermissionDto `json:"permissionBo"`
}
