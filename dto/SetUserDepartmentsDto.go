package dto

type SetUserDepartmentsDto struct {
	Departments []SetUserDepartmentDto `json:"departments"`
	UserId      string                 `json:"userId"`
}
