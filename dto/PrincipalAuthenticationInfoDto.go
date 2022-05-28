package dto


type PrincipalAuthenticationInfoDto struct{
    Authenticated  bool `json:"authenticated"`
    PrincipalType  string `json:"principalType"`
    PrincipalCode  string `json:"principalCode"`
    PrincipalName  string `json:"principalName"`
    AuthenticatedAt  string `json:"authenticatedAt"`
}

