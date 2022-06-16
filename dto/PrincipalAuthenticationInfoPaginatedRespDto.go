package dto

type PrincipalAuthenticationInfoPaginatedRespDto struct {
	StatusCode int                                  `json:"statusCode"`
	Message    string                               `json:"message"`
	ApiCode    int                                  `json:"apiCode,omitempty"`
	Data       PrincipalAuthenticationInfoPagingDto `json:"data"`
}
