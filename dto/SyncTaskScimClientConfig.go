package dto


type SyncTaskScimClientConfig struct{
    OrgUrl  string `json:"org_url,omitempty"`
    UserUrl  string `json:"user_url"`
    Token  string `json:"token"`
    RootDepartmentId  string `json:"root_department_id,omitempty"`
    ParentDepartment  string `json:"parent_department,omitempty"`
    Department  string `json:"department,omitempty"`
}

