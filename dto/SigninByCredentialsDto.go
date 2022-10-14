package dto


type SigninByCredentialsDto struct{
    Connection  string  `json:"connection"`
    PasswordPayload  AuthenticateByPasswordDto `json:"passwordPayload,omitempty"`
    PassCodePayload  AuthenticateByPassCodeDto `json:"passCodePayload,omitempty"`
    AdPayload  AuthenticateByADDto `json:"adPayload,omitempty"`
    LdapPayload  AuthenticateByLDAPDto `json:"ldapPayload,omitempty"`
    Options  SignInOptionsDto `json:"options,omitempty"`
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
}

