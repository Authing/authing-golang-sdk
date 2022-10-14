package dto


type ExtIdpInfoDto struct{
    Identifier  string `json:"identifier"`
    ExtIdpId  string `json:"extIdpId"`
    Type  string  `json:"type"`
    ExtIdpType  string  `json:"extIdpType"`
    BindUrl  string `json:"bindUrl"`
    Name  string `json:"name"`
    NameEn  string `json:"name_en,omitempty"`
    Desc  string `json:"desc,omitempty"`
    DescEn  string `json:"desc_en,omitempty"`
    Logo  string `json:"logo,omitempty"`
}

