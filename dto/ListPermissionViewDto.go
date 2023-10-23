package dto


type ListPermissionViewDto struct{
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
    Keyword  string `json:"keyword,omitempty"`
}

