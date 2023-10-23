package dto


type SignInByPushOptionsDto struct{
    Scope  string `json:"scope,omitempty"`
    Context  string `json:"context,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

