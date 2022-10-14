package dto


type SignupDto struct{
    Connection  string  `json:"connection"`
    PasswordPayload  SignUpByPasswordDto `json:"passwordPayload,omitempty"`
    PassCodePayload  SignUpByPassCodeDto `json:"passCodePayload,omitempty"`
    Profile  SignupProfileDto `json:"profile,omitempty"`
    Options  SignupOptionsDto `json:"options,omitempty"`
}

