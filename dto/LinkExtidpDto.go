package dto


type LinkExtidpDto struct{
    ExtIdpConnIdentifier string `json:"ext_idp_conn_identifier,omitempty"`
    AppId string `json:"app_id,omitempty"`
    IdToken string `json:"id_token,omitempty"`
}

