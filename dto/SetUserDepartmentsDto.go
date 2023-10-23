package dto


type SetUserDepartmentsDto struct{
    UserId  string `json:"userId"`
    Departments  []SetUserDepartmentDto `json:"departments"`
    Options  SetUserDepartmentsOptionsDto `json:"options,omitempty"`
}

