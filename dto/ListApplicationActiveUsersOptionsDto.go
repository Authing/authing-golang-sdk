package dto


type ListApplicationActiveUsersOptionsDto struct{
    Pagination  PaginationDto `json:"pagination,omitempty"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
    WithIdentities  bool `json:"withIdentities,omitempty"`
    WithDepartmentIds  bool `json:"withDepartmentIds,omitempty"`
}

