package dto


type SignInByPasswordPayloadDto struct{
    Password  string `json:"password"`
    Account  string `json:"account,omitempty"`
    Email  string `json:"email,omitempty"`
    Username  string `json:"username,omitempty"`
    Phone  string `json:"phone,omitempty"`
}

