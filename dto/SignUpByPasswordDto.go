package dto


type SignUpByPasswordDto struct{
    Password  string `json:"password"`
    Username  string `json:"username,omitempty"`
    Email  string `json:"email,omitempty"`
}

