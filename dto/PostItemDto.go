package dto


type PostItemDto struct{
    Code  string `json:"code"`
    Name  string `json:"name"`
    Description  string `json:"description,omitempty"`
    UserCount  int `json:"userCount"`
    DepartmentCount  int `json:"departmentCount"`
    MetadataSource  []string `json:"metadataSource"`
}

