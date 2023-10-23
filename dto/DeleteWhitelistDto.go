package dto


type DeleteWhitelistDto struct{
    Type  string  `json:"type"`
    List  []string `json:"list,omitempty"`
}

