package dto


type DepartmentPagingDto struct{
    TotalCount  bool `json:"totalCount"`
    List  []DepartmentDto `json:"list"`
}

