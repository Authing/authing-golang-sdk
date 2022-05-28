package dto


type UserDepartmentRespDto struct{
    DepartmentId  string `json:"departmentId"`
    Name  string `json:"name"`
    Description  string `json:"description"`
    IsLeader  bool `json:"isLeader"`
    Code  string `json:"code"`
    IsMainDepartment  bool `json:"isMainDepartment"`
}

