package dto

type CreateGroupBatchReqDto struct {
	List []CreateGroupReqDto `json:"list"`
}
