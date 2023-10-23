package dto


type SignInByWebAuthnOptionsDto struct{
    Scope  string `json:"scope,omitempty"`
    Context  interface{} `json:"context,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

