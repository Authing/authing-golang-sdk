package dto

type OpenResource struct {
	ResourceCode string      `json:"resourceCode"`
	Authorize    interface{} `json:"authorize"`
}
