package dto


type ApplicationEnabledExtIdpConnDto struct{
    IsSocial  bool `json:"isSocial"`
    ExtIdpId  string `json:"extIdpId"`
    ExtIdpType  string  `json:"extIdpType"`
    ExtIdpConnId  string `json:"extIdpConnId"`
    ExtIdpConnType  string  `json:"extIdpConnType"`
    ExtIdpConnIdentifier  string `json:"extIdpConnIdentifier"`
    ExtIdpConnDisplayName  string `json:"extIdpConnDisplayName"`
    ExtIdpConnLogo  string `json:"extIdpConnLogo"`
}

