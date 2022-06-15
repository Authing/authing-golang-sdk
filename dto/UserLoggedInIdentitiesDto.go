package dto

type UserLoggedInIdentitiesDto struct {
	IdentityId string `json:"identityId"`
	IdpName    string `json:"idpName"`
	IdpNameEn  string `json:"idpNameEn"`
	IdpLogo    string `json:"idpLogo"`
}
