package dto


type MfaTokenIntrospectResponse struct{
    Active  bool `json:"active"`
    UserPoolId  string `json:"userPoolId,omitempty"`
    UserId  string `json:"userId,omitempty"`
    Exp  int `json:"exp,omitempty"`
    Iat  int `json:"iat,omitempty"`
}

