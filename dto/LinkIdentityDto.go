package dto


type LinkIdentityDto struct{
    UserIdInIdp  string `json:"userIdInIdp"`
    UserId  string `json:"userId"`
    ExtIdpId  string `json:"extIdpId"`
    Type  string `json:"type,omitempty"`
    IsSocial  bool `json:"isSocial,omitempty"`
}

