package dto


type DepartmentDto struct{
    DepartmentId  string `json:"departmentId"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    ParentDepartmentId  string `json:"parentDepartmentId"`
    ParentOpenDepartmentId  string `json:"parentOpenDepartmentId,omitempty"`
    Name  string `json:"name"`
    Description  string `json:"description"`
    Code  string `json:"code,omitempty"`
    LeaderUserId  string `json:"leaderUserId,omitempty"`
    MembersCount  int `json:"membersCount"`
    HasChildren  bool `json:"hasChildren"`
}

