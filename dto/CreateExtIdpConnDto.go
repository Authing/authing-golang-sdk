package dto


type CreateExtIdpConnDto struct{
    Fields  interface{} `json:"fields"`
    DisplayName  string `json:"displayName"`
    Identifier  string `json:"identifier"`
    Type  string  `json:"type"`
    ExtIdpId  string `json:"extIdpId"`
    LoginOnly  bool `json:"loginOnly,omitempty"`
    Logo  string `json:"logo,omitempty"`
}

