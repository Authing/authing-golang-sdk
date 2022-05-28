package dto


type SetUserDepartmentDto struct{
    DepartmentId  string `json:"departmentId"`
    IsLeader  bool `json:"isLeader,omitempty"`
    IsMainDepartment  bool `json:"isMainDepartment,omitempty"`
}

