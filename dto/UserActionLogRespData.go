package dto


type UserActionLogRespData struct{
    TotalCount  int `json:"totalCount"`
    List  []UserActionLogDto `json:"list"`
}

