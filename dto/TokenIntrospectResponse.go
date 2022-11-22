package dto


type TokenIntrospectResponse struct{
    Active  bool `json:"active"`
    Sub  string `json:"sub,omitempty"`
    ClientId  string `json:"client_id,omitempty"`
    Exp  int `json:"exp,omitempty"`
    Iat  int `json:"iat,omitempty"`
    Iss  string `json:"iss,omitempty"`
    Jti  string `json:"jti,omitempty"`
    Scope  string `json:"scope,omitempty"`
    TokenType  string `json:"token_type,omitempty"`
}

