package dto


type ApplicationDefaultLoginMethodInput struct{
    ConnectionType  string  `json:"connectionType"`
    QrcodeType  string  `json:"qrcodeType,omitempty"`
    QrcodeExtIdpConnId  string `json:"qrcodeExtIdpConnId,omitempty"`
    AdExtIdpConnId  string `json:"adExtIdpConnId,omitempty"`
    LdapExtIdpConnId  string `json:"ldapExtIdpConnId,omitempty"`
}

