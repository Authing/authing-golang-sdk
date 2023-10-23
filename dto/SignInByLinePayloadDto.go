package dto


type SignInByLinePayloadDto struct{
    AccessToken  string `json:"access_token"`
    IdToken  string `json:"id_token,omitempty"`
}

