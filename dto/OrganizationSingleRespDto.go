package dto

type OrganizationSingleRespDto struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	ApiCode    int             `json:"apiCode,omitempty"`
	Data       OrganizationDto `json:"data"`
}
