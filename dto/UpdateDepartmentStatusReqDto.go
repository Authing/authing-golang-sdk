package dto


type UpdateDepartmentStatusReqDto struct{
    Status  bool `json:"status"`
    DepartmentId  string `json:"departmentId"`
}

