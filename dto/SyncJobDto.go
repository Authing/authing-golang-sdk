package dto


type SyncJobDto struct{
    SyncJobId  int `json:"syncJobId"`
    SyncTaskId  int `json:"syncTaskId"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    SyncStatus  string  `json:"syncStatus"`
    SyncFlow  string  `json:"syncFlow"`
    SyncTrigger  string  `json:"syncTrigger"`
    DepartmentCountAll  int `json:"departmentCountAll"`
    DepartmentCountSucc  int `json:"departmentCountSucc"`
    DepartmentUpdateCountAll  int `json:"departmentUpdateCountAll"`
    DepartmentUpdateCountSucc  int `json:"departmentUpdateCountSucc"`
    AccountCountAll  int `json:"accountCountAll"`
    AccountCountSucc  int `json:"accountCountSucc"`
    AccountUpdateCountAll  int `json:"accountUpdateCountAll"`
    AccountUpdateCountSucc  int `json:"accountUpdateCountSucc"`
    ErrMsg  string `json:"errMsg,omitempty"`
}

