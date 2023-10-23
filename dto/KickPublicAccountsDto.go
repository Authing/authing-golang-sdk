package dto


type KickPublicAccountsDto struct{
    AppIds  []string `json:"appIds"`
    UserId  string `json:"userId"`
    Options  KickPublicAccountsOptionsDto `json:"options,omitempty"`
}

