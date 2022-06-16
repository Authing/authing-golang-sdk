package dto

type AuthorizedResourceDto struct {
	ResourceCode  string            `json:"resourceCode"`
	Description   string            `json:"description,omitempty"`
	Condition     []PolicyCondition `json:"condition,omitempty"`
	ResourceType  string            `json:"resourceType"`
	ApiIdentifier string            `json:"apiIdentifier"`
	Actions       []string          `json:"actions"`
	Effect        string            `json:"effect"`
}
