package dto


type ArchivedUsersListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []ListArchivedUsersRespDto `json:"list"`
}

