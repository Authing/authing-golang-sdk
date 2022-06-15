package dto

type PrincipalAuthenticationInfoPagingDto struct {
	TotalCount int                              `json:"totalCount"`
	List       []PrincipalAuthenticationInfoDto `json:"list"`
}
