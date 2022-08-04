package dto

type ListArchivedUsersRespDto struct {
	UserId     string `json:"userId"`
	ArchivedAt string `json:"archivedAt"`
}
