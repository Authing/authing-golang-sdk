package dto


type UpdateExtIdpConnDto struct{
    Fields  interface{} `json:"fields"`
    DisplayName  string `json:"displayName"`
    Id  string `json:"id"`
    Logo  string `json:"logo,omitempty"`
    LoginOnly  bool `json:"loginOnly,omitempty"`
}

