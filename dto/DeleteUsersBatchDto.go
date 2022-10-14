package dto


type DeleteUsersBatchDto struct{
    UserIds  []string `json:"userIds"`
    Options  DeleteUsersBatchOptionsDto `json:"options,omitempty"`
}

