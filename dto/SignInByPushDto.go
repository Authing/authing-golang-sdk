package dto


type SignInByPushDto struct{
    Account  string `json:"account"`
    Options  SignInByPushOptionsDto `json:"options,omitempty"`
}

