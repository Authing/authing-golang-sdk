package dto


type AddWhitelistDto struct{
    Type  string  `json:"type"`
    List  []string `json:"list,omitempty"`
}

