package dto


type UserListDto struct{
    UserListType string `json:"userListType,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

