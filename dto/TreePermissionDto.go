package dto

type TreePermissionDto struct {
	ResourceId         string                          `json:"resourceId"`
	ResourceType       string                          `json:"resourceType"`
	NodeAuthActionList []DataResourcePolicyTreeStructs `json:"nodeAuthActionList"`
}
