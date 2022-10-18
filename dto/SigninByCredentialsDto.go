package dto


type SigninByCredentialsDto struct{
    Connection  string  `json:"connection"`
    PasswordPayload  SignInByPasswordPayloadDto `json:"passwordPayload,omitempty"`
    PassCodePayload  SignInByPassCodePayloadDto `json:"passCodePayload,omitempty"`
    AdPayload  SignInByAdPayloadDto `json:"adPayload,omitempty"`
    LdapPayload  SignInByLdapPayloadDto `json:"ldapPayload,omitempty"`
    Options  SignInOptionsDto `json:"options,omitempty"`
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
}

