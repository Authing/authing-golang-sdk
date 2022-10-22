package dto


type SignUpDto struct{
    Connection  string  `json:"connection"`
    PasswordPayload  SignUpByPasswordDto `json:"passwordPayload,omitempty"`
    PassCodePayload  SignUpByPassCodeDto `json:"passCodePayload,omitempty"`
    Profile  SignUpProfileDto `json:"profile,omitempty"`
    Options  SignUpOptionsDto `json:"options,omitempty"`
}

