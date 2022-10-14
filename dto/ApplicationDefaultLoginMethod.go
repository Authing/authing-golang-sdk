package dto


type ApplicationDefaultLoginMethod struct{
    ConnectionType  string  `json:"connectionType"`
    QrcodeType  string  `json:"qrcodeType"`
    QrcodeExtIdpConnId  string `json:"qrcodeExtIdpConnId"`
    AdExtIdpConnId  string `json:"adExtIdpConnId"`
    LdapExtIdpConnId  string `json:"ldapExtIdpConnId"`
}

