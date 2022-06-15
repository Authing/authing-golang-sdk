package dto

type AddGroupMembersReqDto struct {
	UserIds []string `json:"userIds"`
	Code    string   `json:"code"`
}
