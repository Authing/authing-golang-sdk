package dto


type GetLdapSubEntriesDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Dn string `json:"dn,omitempty"`
}

