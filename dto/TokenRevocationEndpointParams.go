package dto


type TokenRevocationEndpointParams struct{
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
    Token  string `json:"token"`
}

