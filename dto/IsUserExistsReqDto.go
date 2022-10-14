package dto


type IsUserExistsReqDto struct{
    Username  string `json:"username,omitempty"`
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    ExternalId  string `json:"externalId,omitempty"`
}

