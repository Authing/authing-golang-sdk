package dto

type RoleAuthorizedResourcesRespDto struct {
	ResourceCode  string   `json:"resourceCode"`
	ResourceType  string   `json:"resourceType"`
	Actions       []string `json:"actions"`
	ApiIdentifier string   `json:"apiIdentifier"`
}
