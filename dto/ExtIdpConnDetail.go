package dto


type ExtIdpConnDetail struct{
    Id  string `json:"id"`
    Type  string  `json:"type"`
    ExtIdpId  string `json:"extIdpId"`
    Logo  string `json:"logo"`
    Identifier  string `json:"identifier,omitempty"`
    DisplayName  string `json:"displayName,omitempty"`
    LoginOnly  bool `json:"loginOnly"`
    AssociationMode  string  `json:"associationMode"`
    ChallengeBindingMethods  []string `json:"challengeBindingMethods"`
}

