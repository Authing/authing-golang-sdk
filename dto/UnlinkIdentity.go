package dto


type UnlinkIdentity struct{
    UserId  string `json:"userId"`
    ExtIdpId  string `json:"extIdpId"`
    Type  string `json:"type,omitempty"`
    IsSocial  bool `json:"isSocial,omitempty"`
}

