package dto

type CheckExternalUserPermissionRespDto struct {
	StatusCode int                                `json:"statusCode"`
	Message    string                             `json:"message"`
	ApiCode    int                                `json:"apiCode,omitempty"`
	Data       CheckExternalUserPermissionDataDto `json:"data"`
}
