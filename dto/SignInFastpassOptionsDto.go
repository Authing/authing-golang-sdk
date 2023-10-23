package dto


type SignInFastpassOptionsDto struct{
    Scope  string `json:"scope,omitempty"`
    Context  interface{} `json:"context,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

