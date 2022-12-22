package dto

type GetDataResourceRespDto struct {
	ResourceName  string   `json:"resourceName"`
	ResourceCode  string   `json:"resourceCode"`
	Type          string   `json:"type"`
	Description   string   `json:"description,omitempty"`
	Struct        any      `json:"struct"`
	NamespaceCode string   `json:"namespaceCode"`
	Actions       []string `json:"actions"`
}
