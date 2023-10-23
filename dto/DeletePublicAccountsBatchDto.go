package dto


type DeletePublicAccountsBatchDto struct{
    UserIds  []string `json:"userIds"`
    Options  DeletePublicAccountsBatchOptionsDto `json:"options,omitempty"`
}

