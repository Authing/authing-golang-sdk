package dto

type CreateDataResourceRespDto struct {
	ResourceName string   `json:"resourceName"`
	ResourceCode string   `json:"resourceCode"`
	Type         string   `json:"type"`
	Description  string   `json:"description,omitempty"`
	Struct       any      `json:"struct"`
	Actions      []string `json:"actions"`
}
