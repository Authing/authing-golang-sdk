package dto

type SetUserDepartmentsDto struct {
	Departments []SetUserDepartmentDto       `json:"departments"`
	UserId      string                       `json:"userId"`
	Options     SetUserDepartmentsOptionsDto `json:"options,omitempty"`
}
