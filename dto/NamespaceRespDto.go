package dto

type NamespaceRespDto struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	ApiCode    int          `json:"apiCode,omitempty"`
	Data       NamespaceDto `json:"data"`
}
