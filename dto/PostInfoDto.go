package dto


type PostInfoDto struct{
    Code  string `json:"code"`
    Name  string `json:"name"`
    Description  string `json:"description,omitempty"`
    DepartmentIdList  []string `json:"departmentIdList"`
}

