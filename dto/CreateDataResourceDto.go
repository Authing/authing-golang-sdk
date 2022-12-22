package dto

type CreateDataResourceDto struct {
	Actions       []string `json:"actions"`
	Struct        any      `json:"struct"`
	Type          string   `json:"type"`
	ResourceCode  string   `json:"resourceCode"`
	ResourceName  string   `json:"resourceName"`
	NamespaceCode string   `json:"namespaceCode"`
	Description   string   `json:"description,omitempty"`
}
