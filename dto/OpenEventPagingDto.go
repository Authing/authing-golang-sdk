package dto


type OpenEventPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []OpenEventDto `json:"list"`
}

