package dto

type DepartmentDto struct {
	DepartmentId       string `json:"departmentId"`
	OpenDepartmentId   string `json:"openDepartmentId,omitempty"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	ParentDepartmentId string `json:"parentDepartmentId"`
	Code               string `json:"code,omitempty"`
	LeaderUserId       string `json:"leaderUserId,omitempty"`
	MembersCount       int    `json:"membersCount"`
	HasChildren        bool   `json:"hasChildren"`
}
