package dto


type KickUsersDto struct{
    AppIds  []string `json:"appIds"`
    UserId  string `json:"userId"`
    Options  KickUsersOptionsDto `json:"options,omitempty"`
}

