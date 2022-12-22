package dto

type CreateDepartmentTreeReqDto struct {
	Name     string      `json:"name"`
	Children []string    `json:"children,omitempty"`
	Members  UserInfoDto `json:"members,omitempty"`
}
