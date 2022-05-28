package dto


type ListOrganizationsDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

