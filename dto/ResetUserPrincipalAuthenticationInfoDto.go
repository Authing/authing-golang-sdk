package dto

type ResetUserPrincipalAuthenticationInfoDto struct {
	UserId  string                                         `json:"userId"`
	Options ResetUserPrincipalAuthenticationInfoOptionsDto `json:"options,omitempty"`
}
