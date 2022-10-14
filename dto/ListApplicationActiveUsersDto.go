package dto


type ListApplicationActiveUsersDto struct{
    AppId  string `json:"appId"`
    Options  ListApplicationActiveUsersOptionsDto `json:"options,omitempty"`
}

