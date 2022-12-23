package dto

type UpdateDataResourceRespDto struct {
	ResourceName string      `json:"resourceName"`
	ResourceCode string      `json:"resourceCode"`
	Type         string      `json:"type"`
	Description  string      `json:"description,omitempty"`
	Struct       interface{} `json:"struct"`
	Actions      []string    `json:"actions"`
}
